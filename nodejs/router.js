/**
 * The one and only router in this project
 */

// express router
const router = require('express').Router();

// controllers
const urlController = require('./controllers/url');

// routes
router.get('/', urlController.getList);
router.get('/:url', urlController.redir);
router.post('/', urlController.create);
router.patch('/:url', urlController.update);
router.delete('/:url', urlController.delete);

// export the router
module.exports = router;
