const mongoose = require('mongoose')
const Schema = mongoose.Schema

const Contact = new Schema({
  id: { type: String, required: true, index: { unique: true } },
  firstName: { type: String, required: true },
  lastName: { type: String, required: false },
  company: { type: String, required: false },
  phone: { type: Number, required: false },
  ext: { type: Number, required: false },
  cell: { type: Number, required: false },
  email: { type: String, required: true },
  address: { type: String, required: true },
  address2: { type: String, required: true },
  city: { type: String, required: true },
  state: { type: String, required: true },
  zip: { type: Number, required: true }
//   opportunities: [{ type: Schema.Types.ObjectId, ref: 'Opportunity' }]
})

module.exports = mongoose.model('Contact', Contact)
