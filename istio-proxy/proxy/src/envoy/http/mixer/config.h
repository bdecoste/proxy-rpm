/* Copyright 2017 Istio Authors. All Rights Reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *    http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

#pragma once

#include "envoy/json/json_object.h"
#include "mixer/v1/config/client/client_config.pb.h"
#include "src/envoy/utils/config.h"

namespace Envoy {
namespace Http {
namespace Mixer {

// Config for http filter.
class HttpMixerConfig {
 public:
  // Load from envoy filter config in JSON format.
  void Load(const Json::Object& json) {
    Utils::ReadV2Config(json, &http_config_);

    Utils::SetDefaultMixerClusters(http_config_.mutable_transport());
  }

  // The Http client config.
  const ::istio::mixer::v1::config::client::HttpClientConfig& http_config()
      const {
    return http_config_;
  }

  // check cluster
  const std::string& check_cluster() const {
    return http_config_.transport().check_cluster();
  }
  // report cluster
  const std::string& report_cluster() const {
    return http_config_.transport().report_cluster();
  }

 private:
  // The Http client config.
  ::istio::mixer::v1::config::client::HttpClientConfig http_config_;
};

}  // namespace Mixer
}  // namespace Http
}  // namespace Envoy
