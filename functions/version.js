/**
 * Inspects openapi documents to make sure either the server or
 * path contains /<version>/ - but not both.
 */

 import { createRulesetFunction } from '@stoplight/spectral-core';

 var fullPaths = [];
 
 //helper function to build array of all paths
 function buildFullPaths(servers, paths) {
 
   if (servers) {
     for (var i=0; i<servers.length; i++) {
       //skip localhost for development purposes
       if (servers[i].url.match(/http:\/\/localhost/)) {
         continue;
       }
 
       for (const key in paths) {
         //build our full url
         var t = servers[i].url + key;
         fullPaths.push(t)
       }
     }
   } else {
     fullPaths = paths
   }
 
   return fullPaths
 }
 
 export default createRulesetFunction({
     input: null,
     options: {
       type: 'object',
       additionalProperties: false,
       properties: {
        doubleVersionMatcher: true,
       },
       required: ['doubleVersionMatcher'],
     },
   }, (input, options, { path }) => {
   //build an array of our returned values - this lets us check multiple things at once and return them all
   const results = [];
 
   //get values from our passed-in options
   const { doubleVersionMatcher } = options;
 
   //build regex
   var reDoubleVersions = new RegExp(doubleVersionMatcher, 'gi');
 
   //build an array of all the things we need to check
   var fullPaths = buildFullPaths(input.servers, input.paths);
   
   for (const p in fullPaths) {
     //version should be in either the server or path - NOT BOTH
     if (fullPaths[p].match(reDoubleVersions)) {
       results.push({
         message: `${fullPaths[p]} - resources shouldnt include the version in both the server and path`
       });
     }
   }
 
   return results;
 });