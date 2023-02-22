# PacSprite
Questa utility serve per creare il file pac partendo dalle singole carte in formato png.
Il file delle carte in formato pac serve per i gioci di carte che usano la libreria SDL (Invido, Solitario e Tressette).
Per prima cosa mi serve un immagine unica per tutto il mazzo in formato png.
Ogni colonna ha tutte le carte di un segno con il risultato di avere 4 colonne (bastoni, denari coppe e spade).
Per il mazzo del tarocco piemontese, che è l'unico mazzo che non ho in formato pac, sono 14 file.

## Creare il mazzo unico png
Uso ImageMagick a linea di comando in WSL2. Per creare la fila verticale uso:

    convert -append *.png denari.png
Che crea una fila di tutti i denari nel file denari.png. Questo comando va lanciato 
nella directory dei denari. Nota il -append dove il - serve per indicare la direzione dell'append,
in questo caso verticale. Nella directory _tutti_ ho messo le quattro strisce che combino in un
unica immagine all.png con (nota il +  per l'append in orizzontale):

    convert +append *.png all.png

## Creare il Pac
Ora che ho un immagine unica, non mi rimane altro che aggiungere l'header che mi manca 
per avere un formato pac che funzioni nel solitario.

## Formato pac
Vediamo il formato del file mazzo_piac.pac. Si parte con 100 bytes di descrizione. A seguire il timestamp.
I quattro bytes del timestamp sono in LE32 (prima si scrive il byte meno significativo e poi via in crescendo) che nel file pac sono  0x72 0x9b 0x48 0x3f ossia il numero 0x3f489b72, una data del 2003.
Poi c'è il numero delle animazioni in un byte 0x01.
Poi seguono i bytes  0x50  0x01 che sono 0x0150 che è il width dell'immagine combinata con tutte le carte espresso in due byte sempre in LE32. Dopo segue l'altezza 0xa0 0x05, vale a dire 0x05a0. Poi altri due bytes per i frames 0x01 0x00 che mi da 0x0001. Poi siccome c'è un frame, vengono letti 2 bytes per frame.
In questo caso sono 0xff 0xff. Poi parte l'immagine png che continua fino alla fine del file.