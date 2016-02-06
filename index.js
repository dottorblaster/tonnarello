'use strict';

const Hapi = require('hapi'),
  config = require('./config.js'),
  Path = require('path'),
  Inert = require('inert'),
  server = new Hapi.Server({
		connections: {
			routes: {
				files: {
					relativeTo: Path.join(__dirname, '.')
				}
			}
		}
	});

server.connection({
  host: config.host || 'localhost',
  port: config.port || '3000'
});
server.register(Inert, function () {});

server.route({
	method: 'GET',
	path: '/{param*}',
	handler: {
		directory: {
			path: 'public',
			defaultExtension: 'html',
			redirectToSlash: true,
			listing: false,
			index: true
		}
	}
});

server.route({
  method: 'POST',
  path: '/api/new',
  handler: function(req, reply) {
    reply('Hello!');
  }
});

server.route({
  method: 'POST',
  path: '/api/clone',
  handler: function(req, reply) {
    reply('Hello!');
  }
});

server.start(function (err) {
	if (err) {
		console.error('An error occurred: ' + err);
	} else {
		console.log('Tonnarello running at:', server.info.uri);
	}
});

exports.server = server;
