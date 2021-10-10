<template>
  <el-form
      ref="ruleForm"
      :model="ruleForm"
      status-icon
      :rules="rules"
      label-width="120px"
      class="demo-ruleForm"
  >
    <el-form-item label="Name" prop="name">
      <el-input v-model.number="ruleForm.name"></el-input>
    </el-form-item>
    <el-form-item label="Description" prop="description">
      <el-input v-model.number="ruleForm.description"></el-input>
    </el-form-item>
    <el-form-item label="Country" prop="country">
      <el-input v-model.number="ruleForm.country"></el-input>
    </el-form-item>
    <el-form-item label="City" prop="city">
      <el-input v-model.number="ruleForm.city"></el-input>
    </el-form-item>
    <el-form-item label="Address" prop="address">
      <el-input v-model.number="ruleForm.address"></el-input>
    </el-form-item>
    <el-form-item>
      <el-button type="primary" @click="submitForm('ruleForm')"
      >Add Hotel</el-button
      >
      <el-button @click="resetForm('ruleForm')">Reset</el-button>
    </el-form-item>
  </el-form>
</template>

<script>
import {ElNotification} from "element-plus";

export default {
  name: "HotelForm",
  data() {
    const validateName = (rule, value, callback) => {
      if (!value) {
        return callback(new Error('Please input the name'))
      }
      setTimeout(() => {
        if (value.length > 250) {
          callback(new Error('Name length must be less than 250 characters'))
        } else {
          callback()
        }
      }, 1000)
    }
    const validateDescription = (rule, value, callback) => {
      if (value === '') {
        callback(new Error('Please input the description'))
      } else {
        if (value.length > 1000) {
          callback(new Error('Description must be less than 1000 characters'))
        } else {
          callback()
        }
      }
    }
    const validateCountry = (rule, value, callback) => {
      if (value === '') {
        callback(new Error('Please input the country'))
      } else {
        if (value.length > 100) {
          callback(new Error("Country name must be less than 100 characters"))
        } else {
          callback()
        }
      }
    }
    const validateCity = (rule, value, callback) => {
      if (value === '') {
        callback(new Error('Please input the city'))
      } else {
        if (value.length > 100) {
          callback(new Error("City name must be less than 100 characters"))
        } else {
          callback()
        }
      }
    }
    const validateAddress = (rule, value, callback) => {
      if (value === '') {
        callback(new Error('Please input the address'))
      } else {
        if (value.length > 100) {
          callback(new Error("Address must be less than 250 characters"))
        } else {
          callback()
        }
      }
    }
    return {
      ruleForm: {
        name: '',
        description: '',
        country: '',
        city: '',
        address: '',
      },
      rules: {
        name: [{ validator: validateName, trigger: 'blur' }],
        description: [{ validator: validateDescription, trigger: 'blur' }],
        country: [{ validator: validateCountry, trigger: 'blur' }],
        city: [{ validator: validateCity, trigger: 'blur' }],
        address: [{ validator: validateAddress, trigger: 'blur' }],
      },
    }
  },
  methods: {
    submitForm(formName) {
      this.$refs[formName].validate((valid) => {
        if (valid) {
          this.gatewayClient.createHotel({
            hotel: {
              name: this.ruleForm.name,
              description: this.ruleForm.description,
              country: this.ruleForm.description,
              city: this.ruleForm.city,
              address: this.ruleForm.address,
            },
            token: this.userSingletone.token
          }, function(error) {
            if (error) {
              ElNotification({
                title: 'Error',
                message: error.message,
                type: 'error',
              })
            } else {
              ElNotification({
                title: 'Success',
                message: 'Hotel were created',
                type: 'success',
              })
            }
          })
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