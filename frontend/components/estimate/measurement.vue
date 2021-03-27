<template>
  <div>
    <div class="w-full flex jusify-center items-center pt-4 py-8">
      <div>
        <form @submit.stop.prevent="calculateArea">
          <input
            v-model="width"
            type="number"
            class="   "
            placeholder="Width"
          >
          <input v-model="lenght" type="number" placeholder="Lenght">
          <button name="button">
            Add
          </button>
        </form>
      </div>
    </div>
    <!-- Display all measurements-->
    <div v-if="listOfMeasurements.length > 0" class="w-full flex">
      <table class="table-fixed  border border-gray-400">
        <thead>
          <tr>
            <th class="w-56 px-4 py-2 border border-gray-400 bg-gray-200">
              Width
            </th>
            <th class="w-56 px-4 py-2 border border-gray-400 bg-gray-200">
              Lenght
            </th>
            <th class=" w-56 px-4 py-2 border border-gray-400 bg-gray-200">
              SQFT
            </th>
            <th class="w-56 px-4 py-2 border border-gray-400 bg-gray-200" />
          </tr>
        </thead>
        <tbody>
          <tr v-for="(list, index) in listOfMeasurements" :key="index">
            <th class="border border-gray-400 px-4 py-2 ">
              {{ list.width }}
            </th>
            <th class="border border-gray-400 px-4 py-2">
              {{ list.lenght }}
            </th>
            <th class="border border-gray-400 px-4 py-2">
              {{ list.area }}
            </th>
            <th class="border border-gray-400 px-4 py-2">
              <button :data-delete_id="index" @click="deleteMeasurement">
                X
              </button>
            </th>
          </tr>
        </tbody>
        <tfoot>
          <tr>
            <th class="border border-gray-400 px-4 py-2" />
            <th class="border border-gray-400 px-4 py-2">
              Total sqft
            </th>
            <th class="border border-gray-400 px-4 py-2 bg-green-200">
              {{ totalSQFT }}
            </th>
            <th class="px-4 py-2" />
          </tr>
        </tfoot>
      </table>
    </div>
  </div>
</template>

<script>
export default {
  data () {
    return {
      width: 0,
      lenght: 0,
      area: 0,
      totalSQFT: 0,
      isCalOpen: false,
      listOfMeasurements: []
    }
  },
  computed: {
    islistOfMeasurements () {
      return this.listOfMeasurements > 0
    }
  },
  methods: {
    calculateArea () {
      this.area = parseFloat(this.width) * parseFloat(this.lenght)
      this.listOfMeasurements.push({
        width: parseFloat(this.width),
        lenght: parseFloat(this.lenght),
        area: this.area
      })
      this.newMeasurementCal()
      this.isCalOpen = false
      this.width = 0
      this.lenght = 0
      this.area = 0
    },
    newMeasurementCal () {
      this.totalSQFT = 0
      this.listOfMeasurements.forEach((obj) => {
        this.totalSQFT += obj.area
      })
    },
    deleteMeasurement (event) {
      const indexNumber = event.target.dataset.delete_id

      if (typeof this.listOfMeasurements[indexNumber] !== 'undefined') {
        const deleteMeasurement = this.listOfMeasurements[indexNumber]
        const index = this.listOfMeasurements.indexOf(deleteMeasurement)
        if (index > -1) {
          this.listOfMeasurements.splice(index, 1)
        }
        this.newMeasurementCal()
      }
    }
  }

}
</script>

<style>

</style>
