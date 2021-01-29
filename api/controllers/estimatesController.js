const validator = require('express-validator')
const Estimate = require('../models/Estimate')

// Get all
module.exports.list = function (req, res, next) {
  Estimate.find({}, function (err, estimates) {
    if (err) {
      return res.status(500).json({
        message: 'Error getting records.'
      })
    }
    return res.json(estimates)
  })
}

// Get one
module.exports.show = function (req, res) {
  const id = req.params.id
  Estimate.findOne({ _id: id }, function (err, estimate) {
    if (err) {
      return res.status(500).json({
        message: 'Error getting record.'
      })
    }
    if (!estimate) {
      return res.status(404).json({
        message: 'No such record'
      })
    }
    return res.json(estimate)
  })
}

// Create
module.exports.create = [
  // validations rules
  validator.body('title', 'Please enter Estimate Title').isLength({ min: 1 }),
  validator.body('title').custom((value) => {
    return Estimate.findOne({ title: value }).then((estimate) => {
      if (estimate !== null) {
        return Promise.reject(new Error('Title already in use'))
      }
    })
  }),
  validator.body('author', 'Please enter Author Name').isLength({ min: 1 }),
  validator.body('body', 'Please enter Estimate Content').isLength({ min: 1 }),

  function (req, res) {
    // throw validation errors
    const errors = validator.validationResult(req)
    if (!errors.isEmpty()) {
      return res.status(422).json({ errors: errors.mapped() })
    }

    // initialize record
    const estimate = new Estimate({
      title: req.body.title,
      author: req.body.author,
      body: req.body.body
    })

    // save record
    estimate.save(function (err, estimate) {
      if (err) {
        return res.status(500).json({
          message: 'Error saving record',
          error: err
        })
      }
      return res.json({
        message: 'saved',
        _id: estimate._id
      })
    })
  }
]

// Update
module.exports.update = [
  // validation rules
  validator.body('title', 'Please enter Estimate Title').isLength({ min: 1 }),
  validator.body('title').custom((value, { req }) => {
    return Estimate.findOne({ title: value, _id: { $ne: req.params.id } })
      .then((estimate) => {
        if (estimate !== null) {
          return Promise.reject(new Error('Title already in use'))
        }
      })
  }),
  validator.body('author', 'Please enter Author Name').isLength({ min: 1 }),
  validator.body('body', 'Please enter Estimate Content').isLength({ min: 1 }),

  function (req, res) {
    // throw validation errors
    const errors = validator.validationResult(req)
    if (!errors.isEmpty()) {
      return res.status(422).json({ errors: errors.mapped() })
    }

    const id = req.params.id
    Estimate.findOne({ _id: id }, function (err, estimate) {
      if (err) {
        return res.status(500).json({
          message: 'Error saving record',
          error: err
        })
      }
      if (!estimate) {
        return res.status(404).json({
          message: 'No such record'
        })
      }

      // initialize record
      estimate.title = req.body.title ? req.body.title : estimate.title
      estimate.author = req.body.author ? req.body.author : estimate.author
      estimate.body = req.body.body ? req.body.body : estimate.body

      // save record
      estimate.save(function (err, estimate) {
        if (err) {
          return res.status(500).json({
            message: 'Error getting record.'
          })
        }
        if (!estimate) {
          return res.status(404).json({
            message: 'No such record'
          })
        }
        return res.json(estimate)
      })
    })
  }

]

// Delete
module.exports.delete = function (req, res) {
  const id = req.params.id
  Estimate.findByIdAndRemove(id, function (err, estimate) {
    if (err) {
      return res.status(500).json({
        message: 'Error getting record.'
      })
    }
    return res.json(estimate)
  })
}
