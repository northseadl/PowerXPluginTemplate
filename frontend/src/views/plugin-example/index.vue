<template>
  <div>
    <a-input-number v-model="state.id" placeholder="请输入员工id"></a-input-number>
    <a-button type="primary" @click="onClick">查询</a-button>
  </div>
</template>

<script lang="ts" setup>
  import { reactive } from 'vue';
  import { getEmployee } from '@/api/plugin-example';
  import { Message } from '@arco-design/web-vue';

  const state = reactive({
    id: 0,
  });

  const onClick = () => {
    getEmployee(state.id).then((res) => {
      Message.info(JSON.stringify(res.data));
    });
  };
</script>
