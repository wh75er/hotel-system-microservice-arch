<template>
  <div class="date">
    <div class="demonstration">Reservation date</div>
    <el-date-picker v-model="value" value-format="x" type="date" placeholder="Pick a day">
    </el-date-picker>
  </div>
</template>

<script>
import { defineComponent, reactive, toRefs } from 'vue'

export default defineComponent({
  setup() {
    const state = reactive({
      disabledDate(time) {
        return time.getTime() > Date.now()
      },
      value: '',
    })

    return {
      ...toRefs(state),
    }
  },
})
</script>
