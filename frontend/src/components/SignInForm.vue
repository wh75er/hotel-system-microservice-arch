<template>
  <el-form
      ref="ruleForm"
      :model="ruleForm"
      status-icon
      :rules="rules"
      label-width="120px"
      class="demo-ruleForm"
  >
    <el-form-item label="Login" prop="login">
      <el-input v-model.number="ruleForm.login"></el-input>
    </el-form-item>
    <el-form-item label="Password" prop="pass">
      <el-input
          v-model="ruleForm.pass"
          type="password"
          autocomplete="off"
      ></el-input>
    </el-form-item>
    <el-form-item>
      <el-button type="primary" @click="submitForm('ruleForm')"
      >Log In</el-button
      >
      <el-button @click="resetForm('ruleForm')">Reset</el-button>
    </el-form-item>
  </el-form>
</template>

<script>
import { ElNotification } from 'element-plus'
import Events from '../consts/events.js'

export default {
  mounted() {
  },
  data() {
    const checkLogin = (rule, value, callback) => {
      if (!value) {
        return callback(new Error('Please input the login'))
      } else {
        callback()
      }
    }
    const validatePass = (rule, value, callback) => {
      if (value === '') {
        callback(new Error('Please input the password'))
      } else {
        callback()
      }
    }
    return {
      ruleForm: {
        pass: '',
        login: '',
      },
      rules: {
        pass: [{ validator: validatePass, trigger: 'blur' }],
        login: [{ validator: checkLogin, trigger: 'blur' }],
      },
    }
  },
  methods: {
    submitForm(formName) {
      this.$refs[formName].validate((valid) => {
        if (valid) {
          this.gatewayClient.login({
            login: this.ruleForm.login,
            password: this.ruleForm.pass
          }, function(error, response) {
            if (error) {
              ElNotification({
                title: 'Error',
                message: error.message,
                type: 'error',
              })
            } else {
              ElNotification({
                title: 'Success',
                message: 'You logged in',
                type: 'success',
              })
              const token = response.getValue()
              console.log('token: ', token)
              this.emitter.emit(Events.userLoggedIn, token)
            }
          }.bind(this))
        } else {
          ElNotification({
            title: 'Error',
            message: 'Failed to submit',
            type: 'error',
          })
          return false
        }
      })
    },
    resetForm(formName) {
      this.$refs[formName].resetFields()
    },
  },
}
</script>

<style>
</style>