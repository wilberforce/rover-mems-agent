var mems1xOnlyDivs = document.getElementsByClassName("hideMems1xOnly");
var mems19OnlyDivs = document.getElementsByClassName("hideMems19Only");
var mems2jOnlyDivs = document.getElementsByClassName("hideMems2jOnly");
var mems3OnlyDivs = document.getElementsByClassName("hideMems3Only");
var rc5OnlyDivs = document.getElementsByClassName("hideRc5Only");

//var agentAddress = "http://localhost:8080"
let agentAddress = '//'+location.hostname + ':' + location.port;

export { mems19OnlyDivs, mems1xOnlyDivs, mems2jOnlyDivs, mems3OnlyDivs, rc5OnlyDivs, agentAddress };
