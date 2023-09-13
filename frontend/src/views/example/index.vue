<template>
  <div>
    <a-input
      v-model:value="state.employeeId"
      placeholder="请输入员工ID"
      style="width: 200px; margin-right: 16px"
    />
    <a-button type="primary" @click="onSearch">查询</a-button>
  </div>
</template>

<script lang="ts" setup>
  import { reactive } from 'vue';
  import { getEmployeeInfo } from '@/api/example';
  import { Message } from '@arco-design/web-vue';

  const state = reactive({
    employeeId: '',
  });

  const onSearch = () => {
    getEmployeeInfo(state.employeeId).then((res) => {
      Message.success(res.name);
    });
  };
</script>
