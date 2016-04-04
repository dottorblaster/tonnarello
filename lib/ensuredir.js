var Fs = require('fs');

function ensuredir(path) {
	Fs.lstat(path, function(err, stats) {
		if (!(!err && stats.isDirectory())) {
    console.log('Creating directory: ' + path);
			Fs.mkdir(path);
		}
	});
}

module.exports = ensuredir;
