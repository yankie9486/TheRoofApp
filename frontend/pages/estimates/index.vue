<template>
  <div class="container mr-auto ml-auto py-8">
    <div class="w-full flex">
      <h1 class="text-3xl text-gray-700 font-bold font-sans mt-4 mb-8">
        Estimate Overview
      </h1>
    </div>
    <div class="w-full flex">
      <div class="w-1/2 mx-auto rounded-lg shadow-lg p-4" />
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

<style lang="scss">
input,
button {
	@apply border border-gray-400   py-2 px-4 ;
}

fieldset {
  @apply p-4 bg-gray-300 rounded-md border-gray-400 border;
}

legend {
  @apply text-xl text-gray-700;
}
</style>
