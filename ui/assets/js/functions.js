//Put your JS functions in here
$(document).ready(function(){
      console.log('document ready...');

      showTime();

      $("#page-1").click(function() {
            console.log('page-1');
            $('#tabs-1').removeClass("display-none").addClass("display");
            $('#tabs-2').removeClass("display").addClass("display-none");
            $('#tabs-3').removeClass("display").addClass("display-none");
      });

      $("#page-2").click(function() {
            console.log('page-2');
            $('#tabs-2').removeClass("display-none").addClass("display");
            $('#tabs-1').removeClass("display").addClass("display-none");
            $('#tabs-3').removeClass("display").addClass("display-none");
      });
      $("#page-3").click(function() {
            console.log('page-3');
            $('#tabs-3').removeClass("display-none").addClass("display");
            $('#tabs-1').removeClass("display").addClass("display-none");
            $('#tabs-2').removeClass("display").addClass("display-none");
      });
});


var apihostname = document.getElementById("API_HOSTNAME").textContent;
console.log(apihostname);

var table1 = new Tabulator("#table1", {
ajaxResponse: function( url, params, response){
  return response.Trades;
},
ajaxURL: apihostname + "/db",
layout:"fitDataStretch",
responsiveLayout:"collapse",
dataTree:true,
dataTreeStartExpanded:true,
paginationCounter:"rows",
pagination:"local",
paginationSize:30,
columns:[
            {formatter:"responsiveCollapse", minWidth:30, hozAlign:"center", resizable:false, headerSort:false},
            {title:"Date", field:"Date", responsive:0},
            {title:"Type", field:"Type", hozAlign:"left", sorter:"number",headerFilter:"input"},
            {title:"ReceivedQuantity", field:"ReceivedQuantity", hozAlign:"left",bottomCalc:"sum", bottomCalcParams:{precision:2} },
            {title:"ReceivedCurrency", field:"ReceivedCurrency",hozAlign:"left", headerFilter:"input"},
            {title:"SentQuantity", field:"SentQuantity", hozAlign:"left", sorter:"date" ,bottomCalc:"sum", bottomCalcParams:{precision:2}},
            {title:"SentCurrency", field:"SentCurrency", hozAlign:"left", headerFilter:"input"},
            {title:"FeeAmount", field:"FeeAmount", hozAlign:"left", sorter:"number"},
            {title:"FeeCurrency", field:"FeeCurrency", hozAlign:"left"},
      ],
});

function showTime(){
      var date = new Date();
      var h = date.getHours(); // 0 - 23
      var m = date.getMinutes(); // 0 - 59
      var s = date.getSeconds(); // 0 - 59
      var month = date.getMonth();
      var day = date.getDate();
      var session = "AM";
      
      if(h == 0){
          h = 12;
      }
      
      if(h > 12){
          h = h - 12;
          session = "PM";
      }
      
      h = (h < 10) ? "0" + h : h;
      m = (m < 10) ? "0" + m : m;
      s = (s < 10) ? "0" + s : s;
      
      var time = h + ":" + m + ":" + s + " " + session;
      var date = month + "/" + day;

      document.getElementById("time").innerText = time;
      document.getElementById("time").textContent = date + " " + time;
      
      setTimeout(showTime, 1000);
  }
