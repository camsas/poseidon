/*
Copyright 2018 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package test

import (
	"flag"
	"fmt"
	"github.com/golang/glog"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	batchv1 "k8s.io/api/batch/v1"
	"k8s.io/api/core/v1"
	"k8s.io/api/extensions/v1beta1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/poseidon/test/e2e/framework"
	"math/rand"
	"os"
	"time"
)

const TEST_NAMESPACE = "test"

var testKubeConfig = flag.String("testKubeConfig", "/home/ubuntu/.kube/config", "Specify testKubeConfig path eg: /root/kubeconfig")
var clientset *kubernetes.Clientset
var f *framework.Framework

// go test -args --testKubeConfig="/root/admin.conf"

var _ = Describe("Poseidon", func() {
	//var err error
	flag.Parse()
	hostname, _ := os.Hostname()
	glog.Info("Inside Poseidon tests for k8s:", hostname)

	Describe("Add Pod using Poseidon scheduler", func() {
		glog.Info("Inside Check for adding pod using Poseidon scheduler")
		Context("using firmament for configuring pod", func() {
			name := fmt.Sprintf("test-nginx-pod-%d", rand.Uint32())

			It("should succeed deploying pod using firmament scheduler", func() {
				labels := make(map[string]string)
				labels["schedulerName"] = "poseidon"
				//Create a K8s Pod with poseidon
				pod, err := clientset.CoreV1().Pods(TEST_NAMESPACE).Create(&v1.Pod{
					ObjectMeta: metav1.ObjectMeta{
						Name:   name,
						Labels: labels,
					},
					Spec: v1.PodSpec{
						SchedulerName: "poseidon",
						Containers: []v1.Container{{
							Name:            fmt.Sprintf("container-%s", name),
							Image:           "nginx:latest",
							ImagePullPolicy: "IfNotPresent",
						}}},
				})

				Expect(err).NotTo(HaveOccurred())

				By("Waiting for the pod to have running status")
				f.WaitForPodRunning(pod.Name)
				pod, err = clientset.CoreV1().Pods(TEST_NAMESPACE).Get(name, metav1.GetOptions{})
				Expect(err).NotTo(HaveOccurred())
				glog.Info("pod status =", string(pod.Status.Phase))
				Expect(string(pod.Status.Phase)).To(Equal("Running"))

				By("Pod was in Running state... Time to delete the pod now...")
				err = clientset.CoreV1().Pods(TEST_NAMESPACE).Delete(name, &metav1.DeleteOptions{})
				Expect(err).NotTo(HaveOccurred())
				By("Waiting 5 seconds")
				time.Sleep(time.Duration(5 * time.Second))
				By("Check for pod deletion")
				_, err = clientset.CoreV1().Pods(TEST_NAMESPACE).Get(name, metav1.GetOptions{})
				if err != nil {
					Expect(errors.IsNotFound(err)).To(Equal(true))
				}
			})
		})
	})

	Describe("Add Deployment using Poseidon scheduler", func() {
		glog.Info("Inside Check for adding Deployment using Poseidon scheduler")
		Context("using firmament for configuring Deployment", func() {
			name := fmt.Sprintf("test-nginx-deploy-%d", rand.Uint32())

			It("should succeed deploying Deployment using firmament scheduler", func() {
				// Create a K8s Deployment with poseidon scheduler
				var replicas int32
				replicas = 2
				deployment, err := clientset.ExtensionsV1beta1().Deployments(TEST_NAMESPACE).Create(&v1beta1.Deployment{
					ObjectMeta: metav1.ObjectMeta{
						Labels: map[string]string{"app": "nginx"},
						Name:   name,
					},
					Spec: v1beta1.DeploymentSpec{
						Selector: &metav1.LabelSelector{
							MatchLabels: map[string]string{"app": "nginx", "name": "test-dep"},
						},
						Replicas: &replicas,
						Template: v1.PodTemplateSpec{
							ObjectMeta: metav1.ObjectMeta{
								Labels: map[string]string{"name": "test-dep", "app": "nginx", "schedulerName": "poseidon"},
								Name:   name,
							},
							Spec: v1.PodSpec{
								SchedulerName: "poseidon",
								Containers: []v1.Container{
									{
										Name:            fmt.Sprintf("container-%s", name),
										Image:           "nginx:latest",
										ImagePullPolicy: "IfNotPresent",
									},
								},
							},
						},
					},
				})

				Expect(err).NotTo(HaveOccurred())

				By("Waiting for the Deployment to have running status")
				f.WaitForDeploymentComplete(deployment)
				deployment, err = clientset.ExtensionsV1beta1().Deployments(TEST_NAMESPACE).Get(name, metav1.GetOptions{})
				Expect(err).NotTo(HaveOccurred())

				By(fmt.Sprintf("Creation of deployment %q in namespace %q succeeded.  Deleting deployment.", deployment.Name, TEST_NAMESPACE))
				Expect(deployment.Status.Replicas).To(Equal(deployment.Status.AvailableReplicas))

				By("Pod was in Running state... Time to delete the deployment now...")
				err = clientset.ExtensionsV1beta1().Deployments(TEST_NAMESPACE).Delete(name, &metav1.DeleteOptions{})
				Expect(err).NotTo(HaveOccurred())
				By("Waiting 5 seconds")
				time.Sleep(time.Duration(5 * time.Second))
				By("Check for deployment deletion")
				_, err = clientset.ExtensionsV1beta1().Deployments(TEST_NAMESPACE).Get(name, metav1.GetOptions{})
				if err != nil {
					Expect(errors.IsNotFound(err)).To(Equal(true))
				}
			})
		})
	})

	Describe("Add ReplicaSet using Poseidon scheduler", func() {
		glog.Info("Inside Check for adding ReplicaSet using Poseidon scheduler")
		Context("using firmament for configuring ReplicaSet", func() {
			name := fmt.Sprintf("test-nginx-rs-%d", rand.Uint32())

			It("should succeed deploying ReplicaSet using firmament scheduler", func() {
				//Create a K8s ReplicaSet with poseidon scheduler
				var replicas int32
				replicas = 3
				replicaSet, err := clientset.ExtensionsV1beta1().ReplicaSets(TEST_NAMESPACE).Create(&v1beta1.ReplicaSet{
					ObjectMeta: metav1.ObjectMeta{
						Name: name,
					},
					Spec: v1beta1.ReplicaSetSpec{
						Replicas: &replicas,
						Selector: &metav1.LabelSelector{
							MatchLabels: map[string]string{"name": "test-rs"},
						},
						Template: v1.PodTemplateSpec{
							ObjectMeta: metav1.ObjectMeta{
								Labels: map[string]string{"name": "test-rs", "schedulerName": "poseidon"},
								Name:   name,
							},
							Spec: v1.PodSpec{
								SchedulerName: "poseidon",
								Containers: []v1.Container{
									{
										Name:            fmt.Sprintf("container-%s", name),
										Image:           "nginx:latest",
										ImagePullPolicy: "IfNotPresent",
									},
								},
							},
						},
					},
				})

				Expect(err).NotTo(HaveOccurred())

				By("Waiting for the ReplicaSet to have running status")
				//f.WaitForReadyReplicaSet(replicaSet.Name)
				replicaSet, err = clientset.ExtensionsV1beta1().ReplicaSets(TEST_NAMESPACE).Get(replicaSet.Name, metav1.GetOptions{})
				Expect(err).NotTo(HaveOccurred())

				By(fmt.Sprintf("Creation of ReplicaSet %q in namespace %q succeeded.  Deleting ReplicaSet.", replicaSet.Name, TEST_NAMESPACE))
				Expect(replicaSet.Status.Replicas).To(Equal(replicaSet.Status.AvailableReplicas))
				By("Pod was in Running state... Time to delete the ReplicaSet now...")
				err = clientset.ExtensionsV1beta1().ReplicaSets(TEST_NAMESPACE).Delete(replicaSet.Name, &metav1.DeleteOptions{})
				Expect(err).NotTo(HaveOccurred())
				By("Waiting 5 seconds")
				time.Sleep(time.Duration(5 * time.Second))
				By("Check for ReplicaSet deletion")
				_, err = clientset.ExtensionsV1beta1().ReplicaSets(TEST_NAMESPACE).Get(name, metav1.GetOptions{})
				if err != nil {
					Expect(errors.IsNotFound(err)).To(Equal(true))
				}
			})
		})
	})

	Describe("Add Job using Poseidon scheduler", func() {
		glog.Info("Inside Check for adding Job using Poseidon scheduler")
		Context("using firmament for configuring Job", func() {
			name := fmt.Sprintf("test-nginx-job-%d", rand.Uint32())

			It("should succeed deploying Job using firmament scheduler", func() {
				labels := make(map[string]string)
				labels["name"] = "test-job"
				//Create a K8s Job with poseidon scheduler
				var parallelism int32 = 2
				job, err := clientset.BatchV1().Jobs(TEST_NAMESPACE).Create(&batchv1.Job{
					ObjectMeta: metav1.ObjectMeta{
						Name:   name,
						Labels: labels,
					},
					Spec: batchv1.JobSpec{
						Parallelism: &parallelism,
						Completions: &parallelism,
						Template: v1.PodTemplateSpec{
							ObjectMeta: metav1.ObjectMeta{
								Labels: labels,
							},
							Spec: v1.PodSpec{
								SchedulerName: "poseidon",
								Containers: []v1.Container{
									{
										Name:            fmt.Sprintf("container-%s", name),
										Image:           "nginx:latest",
										ImagePullPolicy: "IfNotPresent",
									},
								},
								RestartPolicy: "Never",
							},
						},
					},
				})

				Expect(err).NotTo(HaveOccurred())

				By("Waiting for the Job to have running status")
				f.WaitForAllJobPodsRunning(job.Name, parallelism)

				job, err = clientset.BatchV1().Jobs(TEST_NAMESPACE).Get(name, metav1.GetOptions{})
				Expect(err).NotTo(HaveOccurred())

				By(fmt.Sprintf("Creation of Jobs %q in namespace %q succeeded.  Deleting Job.", job.Name, TEST_NAMESPACE))
				Expect(job.Status.Active).To(Equal(parallelism))

				By("Job was in Running state... Time to delete the Job now...")
				err = clientset.BatchV1().Jobs(TEST_NAMESPACE).Delete(name, &metav1.DeleteOptions{})
				Expect(err).NotTo(HaveOccurred())
				By("Waiting 5 seconds")
				time.Sleep(time.Duration(5 * time.Second))
				By("Check for Job deletion")
				_, err = clientset.BatchV1().Jobs(TEST_NAMESPACE).Get(name, metav1.GetOptions{})
				if err != nil {
					Expect(errors.IsNotFound(err)).To(Equal(true))
				}
				Expect("Success").To(Equal("Success"))
			})
		})
	})

})

var _ = BeforeSuite(func() {
	var config *rest.Config
	var err error
	config, err = clientcmd.BuildConfigFromFlags("", *testKubeConfig)
	if err != nil {
		panic(err)
	}
	clientset, err = kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}
	f = framework.NewDefaultFramework("sched-poseidon", clientset)
	f.Namespace = createNamespace(clientset)
})

var _ = AfterSuite(func() {
	// Delete namespace
	err := clientset.CoreV1().Namespaces().Delete(TEST_NAMESPACE, &metav1.DeleteOptions{})
	// Delete all pods
	err = clientset.CoreV1().Pods(TEST_NAMESPACE).DeleteCollection(&metav1.DeleteOptions{}, metav1.ListOptions{})
	if err != nil {
		panic(err)
	}
})

func createNamespace(clientset *kubernetes.Clientset) *v1.Namespace {
	ns, err := clientset.CoreV1().Namespaces().Create(&v1.Namespace{
		ObjectMeta: metav1.ObjectMeta{Name: TEST_NAMESPACE},
	})
	if err != nil {
		if errors.IsAlreadyExists(err) {
			return nil
		} else {
			Expect(err).ShouldNot(HaveOccurred())
			return nil
		}
	}
	By("Waiting 5 seconds")
	time.Sleep(time.Duration(5 * time.Second))
	ns, err = clientset.CoreV1().Namespaces().Get(TEST_NAMESPACE, metav1.GetOptions{})
	Expect(err).ShouldNot(HaveOccurred())
	Expect(ns.Name).To(Equal(TEST_NAMESPACE))
	return ns
}
