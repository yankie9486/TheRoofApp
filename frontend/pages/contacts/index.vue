<template>
  <div class="container mr-auto ml-auto py-8">
    <div class="w-full flex">
      <table>
        <thead>
          <tr>
            <th>
              First Name
            </th>
          </tr>
          <tr>
            <th>
              Middle Name
            </th>
          </tr>
          <tr>
            <th>
              Company
            </th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="contact in contactList" :key="contact.id">
            <td>
              {{ contact.fname }}
            </td>
            <td>
              {{ contact.mname }}
            </td>
            <td>
              {{ contact.lname }}
            </td>
            <td>
              {{ contact.company }}
            </td>
          </tr>
        </tbody>
      </table>
    </div>
  </div>
</template>

<script>
export default {
  async fetch () {
    this.contactList = await fetch(
      'http://localhost:8020/api/contacts'
    ).then(res => res.json())
  },
  data () {
    return {
      contactList: []

    }
  }
}
</script>
