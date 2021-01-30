const { Router } = require('express')
const config = require('../config')

const router = Router()

// Initialize Controller
const estimatesController = require('../controllers/estimatesController')

// Get All
router.get('/estimates', estimatesController.list)

// Get One
router.get('/estimates/:id', estimatesController.show)

// Create
router.post('/estimates', config.isAuthenticated, estimatesController.create)

// Update
router.put('/estimates/:id', config.isAuthenticated, estimatesController.update)

// Delete
router.delete('/estimates/:id', config.isAuthenticated, estimatesController.delete)

module.exports = router
