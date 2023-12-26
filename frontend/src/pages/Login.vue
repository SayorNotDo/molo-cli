<script setup lang='ts'>
  import { Lock, User } from '@element-plus/icons-vue';
  import { reactive, ref } from 'vue';
  import useUserStore from '../store/user';
  import type { FormInstance } from 'element-plus';
  import { ElMessage } from 'element-plus';

  interface UserInfo {
    username: string,
    password: string,
  }

  const formRef = ref<FormInstance>();
  const userStore = useUserStore();
  const userInfo = reactive<UserInfo>({
    username: '',
    password: ''
  });
  const onSubmit = async (loginForm: FormInstance | undefined) => {
    if (!loginForm) return;
    try {
      await userStore.login(userInfo);
    } catch (err) {
      ElMessage.error((err as Error).message);
    } finally {

    }

  };
</script>

<template>
  <el-container class='container'>
    <el-row justify='center'>
      <el-col>
        <h1 :style='{ fontSize: "64px", textAlign: "center"}'>Authorization</h1>
        <p :style='{ textAlign: "center"}'>Enter your credentials</p>
      </el-col>
    </el-row>
    <el-row justify='center'>
      <el-form
        status-icon
        ref='formRef'
        :model='userInfo'
        :style='{width:"40%"}'
      >
        <el-form-item>
          <el-input
            v-model='userInfo.username'
            :prefix-icon='User'
            size='large'
          />
        </el-form-item>
        <el-form-item>
          <el-input
            v-model='userInfo.password'
            :prefix-icon='Lock'
            size='large'
            show-password />
        </el-form-item>
        <el-form-item>
          <el-button type='primary' long @click='onSubmit(formRef)'>Login</el-button>
          <el-button text>Create an account</el-button>
        </el-form-item>
      </el-form>
    </el-row>
    <el-divider />
    <el-footer :style='{textAlign:"center"}'>
      <template #default>
        Run into trouble?
        <el-link>
          <el-text size='large'>Support</el-text>
        </el-link>
      </template>
    </el-footer>
  </el-container>
</template>

<style lang='less' scoped>
  .container {
    display: flex;
    height: calc(100vh - 16px);
  }
</style>