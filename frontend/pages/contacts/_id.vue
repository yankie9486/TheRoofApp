<template>
  <div class="container mr-auto ml-auto py-8">
    <EstimateForm />
  </div>
</template>

<script>
import EstimateForm from '../../components/Estimate/Form.vue'

export default {
  components: { EstimateForm },
  validate ({ params, error }) {
    return !isNaN(+params.id)
  },
  // async asyncData ({ params, error }) {
  //   return new Promise((resolve, reject) => {

  //   }).catch((e) => {
  //     error(new Error())
  //   })

  //   fetch(`http://localhost:8020/api/contacts/${params.id}`)
  //     .then(function (rep) {
  //       console.log(rep)
  //       return rep.json()
  //     }).then(function (data) {
  //       console.log(data)
  //     })
  // },
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
    // updateContact () {
    //   fetch(`http://localhost:8020/api/contacts/${+params.id}`)
    //     .then(function (rep) {
    //       console.log(rep)
    //       return rep.json()
    //     }).then(function (data) {
    //       console.log(data)
    //     })
    // }
  }
}
</script>
