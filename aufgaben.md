# Aufgaben

## 1) ein neuer Fisch

### a) Füge einen weiteren Fisch hinzu.

Nimm als erstes das gleiche Bild wie eines der anderen beiden Fische.

Tipp: gucke in aquarium.go in die main Funktion.
Füge ein neues Objekt in die Liste ein.

### b) Gib dem Fisch ein anderes Aussehen

Jetzt versuche dem Fisch ein anderes Aussehen zu geben.

Tipp: gucke in global.go . Dort ist eine Liste der Bildernamen. 

## 2) Der Fisch kann hin und her schwimmen

Bisher können die Fische nur in eine Richtung geradeaus schwimmen.
Lasse Deinen Fisch immer bis zum Beckenrand schwimmen und dann wieder zurück.

Tipp: die Fische werden mit einem `Verhalten` `ImmerGeradeaus` gesteuert.
Du findest es in `verhalten.go`.

Programmiere ein neues Verhalten `HinUndHer`. Das meiste kannst Du
von `ImmerGeradeaus` kopieren.

Noch ein Tipp: wenn der Fisch die Richtung ändert, musst Du das Bild umdrehen.
Denn Fische schwimmen nicht rückwärts.
Das kannst Du mit `op.GeoM.Scale` erreichen.
Aber das kannst Du ganz am Schluß machen.
