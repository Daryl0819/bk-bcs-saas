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
from backend.utils.basic import getitems


def get_svc_extended_routes(project_id):
    return {}


def get_svc_access_info(manifest, extended_routes):
    """
    {
        'external': {
            'NodePort': ['node_ip:{node_port}'],
        },
        'internal': {
            'ClusterIP': [':{port} {Protocol}']
        }
    }
    """
    access_info = {'external': {}, 'internal': {}}
    svc_type = getitems(manifest, ['spec', 'type'])
    ports = getitems(manifest, ['spec', 'ports'])
    if ports:
        if svc_type == 'ClusterIP':
            access_info['internal'] = {
                'ClusterIP': [f":{p['port']} {p['protocol']}" for p in ports]
            }
        elif svc_type == 'NodePort':
            access_info['external'] = {
                'NodePort': [f"node_ip:{p['nodePort']}" for p in ports]
            }

    return access_info
