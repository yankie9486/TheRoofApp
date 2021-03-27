<template>
  <div>
    <div class="container mr-auto ml-auto py-8">
      <div class="w-full flex">
        <div class="w-1/2 mx-auto rounded-lg shadow-lg p-4">
          <h1 class="text-3xl text-gray-700 font-bold font-sans mt-4 mb-8">
            Create Estimate
          </h1>
          <form>
            <div class="flex flex-col">
              <input
                v-model="address"
                type="text"
                class="w-full   mr-4 mb-4"
                placeholder="Address"
              >
              <input
                v-model="address2"
                type="text"
                class="w-full     mr-4 mb-4"
                placeholder="Address2"
              >
            </div>
            <div class="flex">
              <input
                v-model="city"
                type="text"
                class="w-1/3     mr-4 mb-4"
                placeholder="City"
              >
              <input
                v-model="state"
                type="text"
                class="w-1/3     mr-4 mb-4"
                placeholder="State"
              >
              <input
                v-model="zip"
                type="w-1/3 text"
                class="    mb-4"
                placeholder="Zip"
              >
            </div>
            <div class="text-right">
              <button name="button">
                Create Estimate
              </button>
            </div>
          </form>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  data () {
    return {
      address: '',
      address2: '',
      city: '',
      state: '',
      zip: ''
    }
  },
  methods: {
    CreateEstimate () {
      fetch(`http://localhost:8020/api/contacts/`,
      {
        "method":"POST",
        "data": {
          "address": this.address,
          "address2": this.address2,
          "city": this.city,
          "state": this.state,
          "zip": this.zip,
        }
      })
        .then(function (rep) {
          console.log(rep)
          return rep.json()
        }).then(function (data) {
          console.log(data)
        })
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
