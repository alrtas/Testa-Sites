# Aplicativo CLI para windows que testa Sites

Aplicativo criado com o intuito de automatizar os testes se está funcionando o não uma aplicação WEB, o App, faz uma requisção HTTP usando o método GET e valida o retoro 200.<br>
Por padrão o aplicativo testa 3x com espaçamento de tempo de 5s entre cada bateria de testes, a quantidade de testes pode ser customizada no App.<br>
Todos os testes são salvos em um arquivo CSV separado por " ; ", também é possivel ver os logs no proprio App.<br>
 
### Requisitos:
* deve possuir o arquivo `sites.txt`com os sites que deseja-se testar, separados por quebra de linha, na mesma pasta do executavel
