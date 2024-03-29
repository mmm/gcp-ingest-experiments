#
# Copyright 2019 Google LLC
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     https://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

resource "google_cloud_run_service" "test_run_service" {
  provider = "google-beta"
  project = var.project
  name     = var.service_name
  location = var.region

  metadata {
    namespace = var.project
  }

  spec {
    containers {
      image = var.service_image
    }
  }

}

resource "null_resource" "test_run_service_iam_binding_allow_all" {
	provisioner "local-exec" {
		command = "gcloud beta run services add-iam-policy-binding ${google_cloud_run_service.test_run_service.name} --member='allUsers' --role='roles/run.invoker'"
	}
}
