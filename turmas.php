<html>
   <head>
      <title>OtimaHora</title>
      <script src="scripts/turmas.js"></script>
      <meta charset="UTF-8">
      <link rel="stylesheet" href="styles/main.css">
   </head>
   <body>
      <h1>Turmas: <?php echo $_GET["curso"]?></h1>
      <div class="content">
         <h2>Adicione ou modifique as turmas</h2>
         <div class="itens" id="turmas">
         </div>
         <div class="foot">
            <input type="text" id="new_turma" placeholder="nome da turma">
            <button id="adicionar_turma" onclick="adcionar_turma()">Adicionar</button>
         </div>
      </div>

      <script>
      </script>
   </body>
</html>
