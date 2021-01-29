 <template>
  <div class="container mr-auto ml-auto py-8">
    <div class="w-full flex">
      <h1 class="text-3xl text-gray-700 font-bold font-sans mt-4 mb-8">
        Estimate
      </h1>
    </div>
    <div class="w-full flex jusify-center items-center pt-4 py-8">
      <button
        class="py-2 px-2 flex justify-center items-center border border-gray-300 rounded mr-8"
        @click="isCalOpen = !isCalOpen"
      >
        <svg width="14" height="14" viewBox="0 0 14 14" fill="none" xmlns="http://www.w3.org/2000/svg">
          <path
            fill-rule="evenodd"
            clip-rule="evenodd"
            d="M7 0C7.26522 0 7.51957 0.105357 7.70711 0.292893C7.89464 0.48043 8 0.734784 8 1V6H13C13.2652 6 13.5196 6.10536 13.7071 6.29289C13.8946 6.48043 14 6.73478 14 7C14 7.26522 13.8946 7.51957 13.7071 7.70711C13.5196 7.89464 13.2652 8 13 8H8V13C8 13.2652 7.89464 13.5196 7.70711 13.7071C7.51957 13.8946 7.26522 14 7 14C6.73478 14 6.48043 13.8946 6.29289 13.7071C6.10536 13.5196 6 13.2652 6 13V8H1C0.734784 8 0.48043 7.89464 0.292893 7.70711C0.105357 7.51957 0 7.26522 0 7C0 6.73478 0.105357 6.48043 0.292893 6.29289C0.48043 6.10536 0.734784 6 1 6H6V1C6 0.734784 6.10536 0.48043 6.29289 0.292893C6.48043 0.105357 6.73478 0 7 0Z"
            fill="#59595B"
          />
        </svg>
      </button>
      <div :class="[isCalOpen ? 'flex' : 'hidden']">
        <form @submit.stop.prevent="calculateArea">
          <input
            v-model="width"
            type="number"
            class="border-2 border-gray-300 py-2 px-4"
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
    <div class="w-full flex">
      <table class="table-auto">
        <thead>
          <tr>
            <th class="px-4 py-2">
              Width
            </th>
            <th class="px-4 py-2">
              Lenght
            </th>
            <th class="px-4 py-2">
              SQFT
            </th>
            <th class="px-4 py-2" />
          </tr>
        </thead>
        <tbody>
          <tr v-for="(list, index) in listOfMeasurements" :key="index">
            <th class="border px-4 py-2">
              {{ list.width }}
            </th>
            <th class="border px-4 py-2">
              {{ list.lenght }}
            </th>
            <th class="border px-4 py-2">
              {{ list.area }}
            </th>
            <th class="border px-4 py-2">
              <button :data-delete_id="index" @click="deleteMeasurement">
                X
              </button>
            </th>
          </tr>
        </tbody>
        <tfoot>
          <tr>
            <th class="border px-4 py-2" />
            <th class="border px-4 py-2">
              Total sqft
            </th>
            <th class="border px-4 py-2">
              {{ totalSQFT }}
            </th>
            <th class="border px-4 py-2" />
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

<style lang="scss">
input,
button {
	@apply border-2 border-gray-300 py-2 px-4;
}
</style>
