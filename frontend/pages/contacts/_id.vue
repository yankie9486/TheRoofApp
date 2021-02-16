<template>
  <div class="container mr-auto ml-auto py-8">
    <div class="w-full flex">
      <form @submit.stop.prevent="updateContact">
        <fieldset>
          <legend>Customer Contact</legend>
          <div class="flex w-full">
            <input
              v-model="fname"
              type="text"
              class="w-1/3   mr-4 mb-4"
              placeholder="First Name"
            >
            <input
              v-model="mname"
              type="text"
              class="w-1/3   mr-4 mb-4"
              placeholder="Middle Name"
            >
            <input
              v-model="lname"
              type="text"
              class="w-1/3   mb-4"
              placeholder="Last Name"
            >
          </div>
          <input
            v-model="company"
            type="text"
            class="  w-full mb-4"
            placeholder="Company"
          >

          <div class="flex">
            <input
              v-model="phone"
              type="text"
              class="w-1/3   mr-4 mb-4"
              placeholder="Phone"
            >
            <input
              v-model="cell"
              type="text"
              class="w-1/3   mr-4 mb-4"
              placeholder="Cell"
            >
            <input
              v-model="email"
              type="text"
              class="w-1/3   mb-4"
              placeholder="Email"
            >
          </div>
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
          <div class="flex justify-end">
            <button name="button">
              Save
            </button>
          </div>
        </fieldset>
      </form>
    </div>
  </div>
</template>

<script>
export default {
  validate ({ params }) {
    return !isNaN(+params.id)
  },
  async asyncData({params, error}) {

  },
  data () {
    return {
      fname: '',
      mname: '',
      lname: '',
      company: '',
      phone: '',
      cell: '',
      email: '',
      address: '',
      address2: '',
      city: '',
      state: '',
      zip: '',
      data: {}
    }
  },
  methods: {
    updateContact () {

      fetch(`http://localhost:8020/api/contacts/${+params.id}`)
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
