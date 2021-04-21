# -*- coding: utf-8 -*-
#
# Tencent is pleased to support the open source community by making 蓝鲸智云PaaS平台社区版 (BlueKing PaaS Community Edition) available.
# Copyright (C) 2017-2019 THL A29 Limited, a Tencent company. All rights reserved.
# Licensed under the MIT License (the "License"); you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://opensource.org/licenses/MIT
#
# Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on
# an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the
# specific language governing permissions and limitations under the License.
#
import logging
from typing import Optional

from backend.resources.cluster.models import CtxCluster

from ..resource import ResourceClient

logger = logging.getLogger(__name__)

PREFERRED_API_VERSION = "autoscaling/v2beta2"


class HPA(ResourceClient):
    kind = "HorizontalPodAutoscaler"

    def __init__(self, ctx_cluster: CtxCluster, api_version: Optional[str] = PREFERRED_API_VERSION):
        super().__init__(ctx_cluster, api_version)
