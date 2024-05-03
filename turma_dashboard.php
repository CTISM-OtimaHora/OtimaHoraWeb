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

         .sim {
            color: #4CAF50; /* Green */
         }

         .talvez {
            color: #FFC107; /* Yellow */
         }

         .nao {
            color: #F44336; /* Red */
         }
      </style>

      <div class="calendar">
         <div class="calendar-header" onclick="toggleStatusCol('0')">Segunda-feira</div>
         <div class="calendar-header" onclick="toggleStatusCol('1')">Terça-feira</div>
         <div class="calendar-header" onclick="toggleStatusCol('2')">Quarta-feira</div>
         <div class="calendar-header" onclick="toggleStatusCol('3')">Quinta-feira</div>
         <div class="calendar-header" onclick="toggleStatusCol('4')">Sexta-feira</div>

         <!-- Day 1 -->
         <div id="0-0" class="calendar-cell sim" onclick="toggleStatus(this)">SIM</div>
         <div id="0-1" class="calendar-cell sim" onclick="toggleStatus(this)">SIM</div>
         <div id="0-2" class="calendar-cell sim" onclick="toggleStatus(this)">SIM</div>
         <div id="0-3" class="calendar-cell sim" onclick="toggleStatus(this)">SIM</div>
         <div id="0-4" class="calendar-cell sim" onclick="toggleStatus(this)">SIM</div>

         <!-- Day 2 -->
         <div id="1-0" class="calendar-cell sim" onclick="toggleStatus(this)">SIM</div>
         <div id="1-1" class="calendar-cell sim" onclick="toggleStatus(this)">SIM</div>
         <div id="1-2" class="calendar-cell sim" onclick="toggleStatus(this)">SIM</div>
         <div id="1-3" class="calendar-cell sim" onclick="toggleStatus(this)">SIM</div>
         <div id="1-4" class="calendar-cell sim" onclick="toggleStatus(this)">SIM</div>

         <!-- Day 3 -->
         <div id="2-0" class="calendar-cell sim" onclick="toggleStatus(this)">SIM</div>
         <div id="2-1" class="calendar-cell sim" onclick="toggleStatus(this)">SIM</div>
         <div id="2-2" class="calendar-cell sim" onclick="toggleStatus(this)">SIM</div>
         <div id="2-3" class="calendar-cell sim" onclick="toggleStatus(this)">SIM</div>
         <div id="2-4" class="calendar-cell sim" onclick="toggleStatus(this)">SIM</div>

         <!-- Day 4 -->
         <div id="3-0" class="calendar-cell sim" onclick="toggleStatus(this)">SIM</div>
         <div id="3-1" class="calendar-cell sim" onclick="toggleStatus(this)">SIM</div>
         <div id="3-2" class="calendar-cell sim" onclick="toggleStatus(this)">SIM</div>
         <div id="3-3" class="calendar-cell sim" onclick="toggleStatus(this)">SIM</div>
         <div id="3-4" class="calendar-cell sim" onclick="toggleStatus(this)">SIM</div>

         <!-- Day 5 -->
         <div id="4-0" class="calendar-cell sim" onclick="toggleStatus(this)">SIM</div>
         <div id="4-1" class="calendar-cell sim" onclick="toggleStatus(this)">SIM</div>
         <div id="4-2" class="calendar-cell sim" onclick="toggleStatus(this)">SIM</div>
         <div id="4-3" class="calendar-cell sim" onclick="toggleStatus(this)">SIM</div>
         <div id="4-4" class="calendar-cell sim" onclick="toggleStatus(this)">SIM</div>
      </div>

      <div style="display: flex; flex-direction: row; justify-content: space-between;">
         <div style="display: flex; flex-direction: column;">
            <h3>Adicionar Professor</h3>  
            <input type="text" id="professor-input" placeholder="nome do professor"/>
            <button onclick="add_professor()">Adicionar</button>
         </div>
         <div style="display: flex; flex-direction: column;">
            <h3>Adicionar Disciplina</h3>  
            <input type="text" id="disciplina-input" placeholder="nome da disciplina"/>
            <button onclick="add_disciplina()">Adicionar</button>
         </div>
      </div>

   </body>
</html>
