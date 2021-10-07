<template>
  <el-table class="room-card" :data="rooms" border style="width: 100%">
    <el-table-column prop="RoomType" label="Room Type" width="180" />
    <el-table-column prop="Beds" label="Beds" width="180" />
    <el-table-column prop="Offers" label="Offers" />
    <el-table-column prop="Price" label="Price" width="180" />
    <el-table-column width="75" >
      <template #default="scope">
        <el-button
          type="text"
          size="small"
          @click.prevent="clicked(scope.$index, rooms)"
        >
          Reserve
        </el-button>
      </template>
    </el-table-column>
  </el-table>
</template>

<script>
export default {
  data() {
    return {
      rooms: [
        {
          RoomType: 'Family',
          Beds: 4,
          Offers: [
              "food included",
              "pool included",
          ],
          Price: 400,
        },
        {
          RoomType: 'Lone wolf',
          Beds: 1,
          Offers: [
            "food included",
          ],
          Price: 150,
        },
      ],
    }
  },
  methods: {
    clicked(index, rows) {
      console.log(index, rows)
      console.log(rows[0])
    }
  }
}
</script>

<style>
.room-card {
  box-shadow: var(--el-box-shadow-light);
}
</style>
