<html>
   <head>
      <script src="scripts/dashboard.js"></script>
      <link rel="stylesheet" href="style/main.css">
      <link rel="stylesheet" href="turma_dashboard.css">
   </head>
   <body>

      <div class="calendar">
         <!-- Weekdays header -->
         <div class="calendar-header" onclick="toggleStatusCol('0')">Segunda-Feira</div>
         <div class="calendar-header" onclick="toggleStatusCol('1')">Terça-Feira</div>
         <div class="calendar-header" onclick="toggleStatusCol('2')">Quarta-Feira</div>
         <div class="calendar-header" onclick="toggleStatusCol('3')">Quinta-Feira</div>
         <div class="calendar-header" onclick="toggleStatusCol('4')">Sexta-Feira</div>

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
