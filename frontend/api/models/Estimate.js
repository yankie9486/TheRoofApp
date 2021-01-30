const mongoose = require('mongoose')
const Schema = mongoose.Schema

const Estimate = new Schema({
  _id: Schema.Types.ObjectId,
  estimateN: { type: String, required: true, index: { unique: true } },
  type: { type: String, required: true },
  material: { type: String, required: true }

})

module.exports = mongoose.model('Estimate', Estimate)
