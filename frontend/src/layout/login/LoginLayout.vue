<script setup lang='ts'>
  import loginImg from '../../assets/images/logo-universal.png';
</script>

<template>
  <el-container class='container'>
    <el-row align='middle'>
      <el-col :span='12'>
        <el-image
          :src="loginImg"
        />
      </el-col>
      <el-col :span='12'>
        <transition>
          <router-view />
        </transition>
      </el-col>
    </el-row>
  </el-container>
</template>

<style lang='less' scoped>
.container {
  display: flex;
  height: calc(100vh - 16px);
}
</style>