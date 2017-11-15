'use strict';

var mask = (function(my) {
  var my = {}

  my.formatValor = function (valor) {
    valor = ("R$ "+valor).replace('.',',')

    if (!valor.includes(',')){
      return valor+',00'
    }else{
      valor = valor.split(',')
      if (valor[1].length == 1){
        valor[1] = valor[1]+"0"
      }
      return valor
    }
  }

  my.desFormatValor = function (valor) {
    valor = valor.replace('R$ ', '')
    return parseFloat(valor.replace(',','.'))
  }

  return my

})(mask)
