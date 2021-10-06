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
    <el-form-item label="Confirm" prop="checkPass">
      <el-input
          v-model="ruleForm.checkPass"
          type="password"
          autocomplete="off"
      ></el-input>
    </el-form-item>
    <el-form-item>
      <el-button type="primary" @click="submitForm('ruleForm')"
      >Submit</el-button
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
      }
      setTimeout(() => {
        console.log(value)
        if (value.length < 6) {
          callback(new Error('Login length must be more than 6 characters'))
        } else {
          if (value.length >= 24) {
            callback(new Error('Login must be less than 24 characters'))
          } else {
            if (!/^[a-zA-Z0-9]+$/.test(value)) {
              callback(new Error('Login must consist of latin characters with optional numbers'))
            } else {
              callback()
            }
          }
        }
      }, 1000)
    }
    const validatePass = (rule, value, callback) => {
      if (value === '') {
        callback(new Error('Please input the password'))
      } else {
        if (value.length < 10) {
          callback(new Error('Password must be more than 10 characters'))
        } else {
          if (value.length >= 128) {
            callback(new Error('Password must be less than 128 characters'))
          } else {
            if (!/^[a-zA-Z0-9./?,'[\]\\!@#$%^&*()]+$/.test(value)) {
              callback(new Error('The password must consist of latin chars with optional numbers or special characters ./\\\\?,\'[]!@#$%^&*()'))
            }
            if (this.ruleForm.checkPass !== '') {
              this.$refs.ruleForm.validateField('checkPass')
            }
            callback()
          }
        }
      }
    }
    const validatePass2 = (rule, value, callback) => {
      if (value === '') {
        callback(new Error('Please input the password again'))
      } else if (value !== this.ruleForm.pass) {
        callback(new Error("Two inputs don't match!"))
      } else {
        callback()
      }
    }
    return {
      ruleForm: {
        pass: '',
        checkPass: '',
        login: '',
      },
      rules: {
        pass: [{ validator: validatePass, trigger: 'blur' }],
        checkPass: [{ validator: validatePass2, trigger: 'blur' }],
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
