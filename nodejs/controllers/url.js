/**
 * Database connection
 */
const db = require('../db').get('gonodejsperf');

/**
 * Create indexes
 */
db.collection('url').createIndex({ url: 1 }, { unique: true });

/**
 * Generate pseudorandom string
 * @param {Number} len Length of the expected string
 */
const randomstring = (len) => {
    let alphabet = 'abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789-_';
    let str = '';
    for (let i = 0; i < len; i++) {
        str += alphabet.charAt(Math.floor(Math.random() * alphabet.length));
    }
    return str;
}

/**
 * Get list of all shorened urls
 */
module.exports.getList = (req, res) => {
    db.collection('url').find({}).project({
        _id: false,
        url: true,
        target: true,
        used: true,
        lastUse: true,
    }).toArray((err, data) => {
        if (err) {
            return res.json({
                status: 'ERROR',
                msg: err.toString(),
                data: null,
            });
        }

        // send ok response
        res.json({
            status: 'OK',
            msg: 'Full URL list successfully retrieved.',
            data: data,
        });
    });
};

/**
 * Do the redirect
 */
module.exports.redir = (req, res) => {
    db.collection('url').findOneAndUpdate({
        url: req.params.url,
    }, {
        $inc: {
            used: 1,
        },
        $set: {
            lastUse: new Date(),
        },
    }, {
        projection: {
            _id: false,
            target: true,
        },
        returnOriginal: false,
    }, (err, data) => {
        if (err || !data.value) {
            return res.redirect('/');
        }

        // redirect to the original url
        res.redirect(data.value.target);
    });
};

/**
 * Create new shortened URL
 * @param req.body.target The target URL
 */
module.exports.create = (req, res) => {
    // create document for save
    let doc = {
        url: randomstring(6),
        target: req.body.target,
        used: 0,
        lastUse: null,
        created: new Date(),
        modified: new Date(),
        password: randomstring(32),
    };

    // insert to database
    db.collection('url').insertOne(doc, (err, data) => {
        if (err) {
            return res.json({
                status: 'ERROR',
                msg: err.toString(),
                data: null,
            });
        }

        // send ok response
        res.json({
            status: 'OK',
            msg: 'URL was shortened.',
            data: data.ops[0],
        });
    });
};

/**
 * Update existing URL's target
 * @param req.body.password Password given on url creation
 * @param req.body.target New target URL
 */
module.exports.update = (req, res) => {
    db.collection('url').findOneAndUpdate({
        url: req.params.url,
        password: req.body.password,
    }, {
        $set: {
            modified: new Date(),
            target: req.body.target,
        },
    }, {
        returnOriginal: false,
    }, (err, data) => {
        if (err) {
            return res.json({
                status: 'ERROR',
                msg: err.toString(),
                data: null,
            });
        }

        // test data (if no = password wrong)
        if (!data.lastErrorObject.updatedExisting) {
            return res.json({
                status: 'ERROR',
                msg: 'The password you entered is not valid.',
                data: null,
            });
        }

        // send ok response
        res.json({
            status: 'OK',
            msg: 'The target of the URL was updated successfully.',
            data: data.value,
        });
    });
};

/**
 * Delete existing shorened URL
 * @param req.params.url The shortened url id to delete
 */
module.exports.delete = (req, res) => {
    db.collection('url').deleteOne({
        url: req.params.url,
    }, (err) => {
        if (err) {
            return res.json({
                status: 'ERROR',
                msg: err.toString(),
                data: null,
            });
        }

        // send ok response
        res.json({
            status: 'OK',
            msg: 'Shortened URL was deleted.',
        });
    });
};
