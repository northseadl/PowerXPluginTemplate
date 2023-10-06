import { AppRouteRecordRaw } from '@/router/routes/types';
import { PLUGIN_LAYOUT } from '@/router/routes/base';

const PluginExample: AppRouteRecordRaw = {
  path: '/plugin-example',
  name: 'PluginExample',
  component: PLUGIN_LAYOUT,
  redirect: '/plugin-example/index',
  meta: {
    locale: '插件示例',
    icon: 'icon-settings',
    requiresAuth: true,
  },
  children: [
    {
      path: '/plugin-example/index',
      name: 'PluginExampleIndex',
      component: () => import('@/views/plugin-example/index.vue'),
      meta: {
        locale: '插件示例',
        requiresAuth: true,
      },
    },
  ],
};

export default PluginExample;
