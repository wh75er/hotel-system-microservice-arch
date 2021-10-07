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
      >Sign In</el-button
      >
      <el-button @click="resetForm('ruleForm')">Reset</el-button>
    </el-form-item>
  </el-form>
</template>

<script>
export default {
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
          alert('submit!')
        } else {
          console.log('error submit!!')
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
