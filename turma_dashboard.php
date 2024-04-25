<html>
   <head>
      <script src="scripts/dashboard.js"></script>
      <link rel="stylesheet" href="style/main.css">
   </head>
   <body>
      <style>
         body {
            font-family: Arial, sans-serif;
            margin: 0;
            padding: 0;
            background: #f0f0f0;
         }

         .calendar {
            display: grid;
            grid-template-columns: repeat(5, 1fr);
            grid-template-rows: auto repeat(5, 1fr);
            gap: 1px;
            background: #ccc;
            padding: 10px;
         }

         .calendar-header {
            text-align: center;
            font-weight: bold;
            background: #e0e0e0;
            padding: 10px;
         }

         .calendar-cell {
            background: #fff;
            border: 1px solid #ccc;
            padding: 10px;
            text-align: center;
            cursor: pointer;
            transition: background 0.2s;
         }

         .yes {
            background: #4CAF50; /* Green */
            color: white;
         }

         .maybe {
            background: #FFC107; /* Yellow */
            color: black;
         }

         .no {
            background: #F44336; /* Red */
            color: white;
         }
      </style>

      <div class="calendar">
         <!-- Weekdays header -->
         <div class="calendar-header" onclick="toggleStatusCol('0')">Monday</div>
         <div class="calendar-header" onclick="toggleStatusCol('1')">Tuesday</div>
         <div class="calendar-header" onclick="toggleStatusCol('2')">Wednesday</div>
         <div class="calendar-header" onclick="toggleStatusCol('3')">Thursday</div>
         <div class="calendar-header" onclick="toggleStatusCol('4')">Friday</div>

         <!-- Week content -->
         <!-- Day 1 -->
         <div id="0-0" class="calendar-cell yes" onclick="toggleStatus(this)">YES</div>
         <div id="0-1" class="calendar-cell yes" onclick="toggleStatus(this)">YES</div>
         <div id="0-2" class="calendar-cell yes" onclick="toggleStatus(this)">YES</div>
         <div id="0-3" class="calendar-cell yes" onclick="toggleStatus(this)">YES</div>
         <div id="0-4" class="calendar-cell yes" onclick="toggleStatus(this)">YES</div>

         <!-- Day 2 -->
         <div id="1-0" class="calendar-cell yes" onclick="toggleStatus(this)">YES</div>
         <div id="1-1" class="calendar-cell yes" onclick="toggleStatus(this)">YES</div>
         <div id="1-2" class="calendar-cell yes" onclick="toggleStatus(this)">YES</div>
         <div id="1-3" class="calendar-cell yes" onclick="toggleStatus(this)">YES</div>
         <div id="1-4" class="calendar-cell yes" onclick="toggleStatus(this)">YES</div>

         <!-- Day 3 -->
         <div id="2-0" class="calendar-cell yes" onclick="toggleStatus(this)">YES</div>
         <div id="2-1" class="calendar-cell yes" onclick="toggleStatus(this)">YES</div>
         <div id="2-2" class="calendar-cell yes" onclick="toggleStatus(this)">YES</div>
         <div id="2-3" class="calendar-cell yes" onclick="toggleStatus(this)">YES</div>
         <div id="2-4" class="calendar-cell yes" onclick="toggleStatus(this)">YES</div>

         <!-- Day 4 -->
         <div id="3-0" class="calendar-cell yes" onclick="toggleStatus(this)">YES</div>
         <div id="3-1" class="calendar-cell yes" onclick="toggleStatus(this)">YES</div>
         <div id="3-2" class="calendar-cell yes" onclick="toggleStatus(this)">YES</div>
         <div id="3-3" class="calendar-cell yes" onclick="toggleStatus(this)">YES</div>
         <div id="3-4" class="calendar-cell yes" onclick="toggleStatus(this)">YES</div>

         <!-- Day 5 -->
         <div id="4-0" class="calendar-cell yes" onclick="toggleStatus(this)">YES</div>
         <div id="4-1" class="calendar-cell yes" onclick="toggleStatus(this)">YES</div>
         <div id="4-2" class="calendar-cell yes" onclick="toggleStatus(this)">YES</div>
         <div id="4-3" class="calendar-cell yes" onclick="toggleStatus(this)">YES</div>
         <div id="4-4" class="calendar-cell yes" onclick="toggleStatus(this)">YES</div>
      </div>
   </body>
</html>
