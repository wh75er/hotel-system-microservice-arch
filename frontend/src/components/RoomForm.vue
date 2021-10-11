<template>
  <el-form
      ref="ruleForm"
      :model="ruleForm"
      status-icon
      :rules="rules"
      label-width="120px"
      class="demo-ruleForm"
  >
    <el-form-item label="Room Type" prop="roomType">
      <el-input v-model.number="ruleForm.roomType"></el-input>
    </el-form-item>
    <el-form-item label="Amount" prop="amount">
      <el-input-number v-model="ruleForm.amount" :min="1" :max="100"/>
    </el-form-item>
    <el-form-item label="Beds amount" prop="beds">
      <el-input-number v-model="ruleForm.beds" :min="1" :max="50"/>
    </el-form-item>
    <el-form-item label="Night Price" prop="nightPrice">
      <el-input v-model.number="ruleForm.nightPrice"></el-input>
    </el-form-item>
    <el-form-item label="Add offer" prop="currOffer">
      <el-input v-model.number="ruleForm.currOffer"></el-input>
      <el-button  @click="addOffer(ruleForm.currOffer)">Add</el-button>
      <el-card class="box-card">
        <div v-for="o in ruleForm.offers" :key="o" class="text text-wrapping item">{{ o }}</div>
      </el-card>
    </el-form-item>
    <el-form-item>
      <el-button type="primary" @click="submitForm('ruleForm')"
      >Add Room</el-button
      >
      <el-button @click="resetForm('ruleForm')">Reset</el-button>
    </el-form-item>
  </el-form>
</template>

<script>
import {ElNotification} from "element-plus";

export default {
  name: "RoomForm",
  data() {
    const validateRoomType = (rule, value, callback) => {
      if (!value) {
        return callback(new Error('Please input the room type'))
      }
      setTimeout(() => {
        if (value.length > 250) {
          callback(new Error('Room Type must be less than 250 characters'))
        } else {
          callback()
        }
      }, 1000)
    }
    const validateAmount = (rule, value, callback) => {
      if (value === '') {
        callback(new Error('Please input the amount of rooms'))
      } else {
        if (value <= 0) {
          callback(new Error('Amount of rooms must be positive'))
        } else {
          callback()
        }
      }
    }
    const validateBeds = (rule, value, callback) => {
      if (value === '') {
        callback(new Error('Please input the number of beds'))
      } else {
        if (value <= 0) {
          callback(new Error("Number of beds must be positive"))
        } else {
          callback()
        }
      }
    }
    const validateNightPrice = (rule, value, callback) => {
      if (value === '') {
        callback(new Error('Please input the night price'))
      } else {
        if (typeof value != 'number') {
          callback(new Error('Night price must be a number only'))
        } else {
          if (value <= 0) {
            callback(new Error("Night price must be positive"))
          } else {
            callback()
          }
        }
      }
    }
    return {
      ruleForm: {
        roomType: '',
        amount: 0,
        beds: 0,
        offers: [],
        currOffer: '',
        nightPrice: null,
      },
      rules: {
        roomType: [{ validator: validateRoomType, trigger: 'blur' }],
        amount: [{ validator: validateAmount, trigger: 'blur' }],
        beds: [{ validator: validateBeds, trigger: 'blur' }],
        nightPrice: [{ validator: validateNightPrice, trigger: 'blur' }],
      },
    }
  },
  methods: {
    submitForm(formName) {
      console.log('HOTEL ID IS : ', this.id)
      this.$refs[formName].validate((valid) => {
        if (valid) {
          const hotelUuid = this.$route.params.id
          this.gatewayClient.createRoom({
            room: {
              hotelUuid: hotelUuid,
              roomType: this.ruleForm.roomType,
              amount: this.ruleForm.amount,
              beds: this.ruleForm.beds,
              offers: this.ruleForm.offers,
              nightPrice: this.ruleForm.nightPrice,
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
                message: 'Room were created',
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
      this.ruleForm.offers = []
    },
    addOffer(offer) {
      this.ruleForm.offers.push(offer)
      console.log('Offer was added: ', this.ruleForm.offers)
    }
  },
}
</script>

<style>

.text {
  font-size: 14px;
}

.item {
  padding: 2px 0;
}

.box-card {
  width: 200px;
}
</style>