'use strict';

const Hapi = require('hapi'),
  config = require('./config.js'),
  server = new Hapi.Server();

server.connection({
  host: config.host || 'localhost',
  port: config.port || '3000'
});

server.route({
  method: 'POST',
  path: '/update'
  handler: function(req, reply) {
    reply('Hello!');
  }
});
