var page = require('webpage').create();
var system = require('system');
var url, selector, output;

url = system.args[1]
//'div.grace.full-field' -- squad,
//'table.stat-table' -- league ranking table
selector = system.args[2]
output = system.args[3]

page.open(url, function() {
  // being the actual size of the headless browser
  page.viewportSize = { width: 1920, height: 1080 };

  var clipRect = page.evaluate(function(sel){
    return document.querySelector(sel).getBoundingClientRect();
  }, selector);

  page.clipRect = {
    top:    clipRect.top,
    left:   clipRect.left,
    width:  clipRect.width,
    height: clipRect.height
  };

  page.render(output);
  phantom.exit();
});