# Aufgaben

## 1) ein neuer Fisch

### a) Füge einen weiteren Fisch hinzu.

Nimm als erstes das gleiche Bild wie eines der anderen beiden Fische.

Tipp: gucke in `aquarium.go` in die `main` Funktion.
Füge ein neues Objekt in die Liste ein.

### b) Gib dem Fisch ein anderes Aussehen

Jetzt versuche dem Fisch ein anderes Aussehen zu geben.

Tipp: gucke in `global.go` . Dort ist die Liste der Bildernamen. Füge
Den Bild hinzu.

## 2) Der Fisch kann hin und her schwimmen

Bisher können die Fische nur in eine Richtung geradeaus schwimmen.
Lasse Deinen Fisch immer bis zum Beckenrand schwimmen und dann wieder zurück.

Tipp: die Fische werden mit einem `Verhalten` `ImmerGeradeaus` gesteuert.
Du findest es in `verhalten.go`.

Programmiere ein neues Verhalten `HinUndHer`. Das meiste kannst Du
von `ImmerGeradeaus` kopieren.

Noch ein Tipp: wenn der Fisch die Richtung ändert, musst Du das Bild umdrehen,
denn Fische schwimmen nicht rückwärts.
Das kannst Du mit `op.GeoM.Scale` und einer negativen Zahl erreichen.

## 3) Hintergrund

Bisher ist alles schwarz. Lasse das Programm einen schönen Hintergrund malen.

Tipps:
- `ebiten.Image`, der Typ mit dem wir unser Aquarium zeichnen, hat eine
Methode `Fill`, die alles mit einer gewünschten Farbe füllt.
-  Statt `Figur` zu benutzen, erzeuge einen neuen Typ, der sowohl 
   `BildElement` als auch `Maler` implementiert. Das kannst Du Dir in
   `FrameCounter` in `framecounter.go` abschauen.
   Die `ZeitVergangen` Methode bleibt dabei leer, es sei denn Du möchtest
   zum einfarbigen, statischen Hintergrund noch dynamische Effekte hinzufügen.
   Den neuen Typ mit Methoden packst Du am besten in eine neue Datei 
   `hintergrund.go`.
- Denk daran, dass der Hintergrund zuerst gemalt werden muß. 

## 4) Lebendige Fische

Die Fische sehen noch ziemlich steif aus. Sie schwimmen zwar durch das Wasser,
aber ohne sich auch nur im geringsten zu verändern. Erstelle eine Animation, die
endlos abgespielt wird.

Tipp:
- Zuerst mußt Du Bilder erstellen, die Du wie ein Daumenkino abspielen kannst.
- Dann mußt Du in `ZeitVergangen` nicht nur den Ort ändern, sondern auch das
  benutzte Bild.  
- Du kannst die Bilder abwechselnd vorwärts und rückwärts abspielen, z.B.
  1, 2, 3, 2, 1, 2, 3, ...
- Es gibt inzwischen zwei Implementierungen von `Verhalten`, und die neue
  Funktion wird in beiden benötigt. Du kannst den Code natürlich kopieren und
  in beiden separat einfügen. Evt. fällt Dir hier eine bessere Lösung ein.
- Wahrscheinlich sieht das Ergebnis sehr hippelig aus, weil deine Animation
  vielleicht nur 2-4 Bilder umfaßt, das Bild aber jede 1/60s geändert wird.
  Es wird notwendig sein zu beachten, wieviel Zeit seit dem letzten Aufruf von
  `ZeitVergangen` verstrichen ist, und nicht jedes Mal das Bild zu ändern.
  
## 5) Langsame und schnelle Fische

Die Fische bewegen sich und sind hübsch animiert, aber es gibt noch ein Problem.
Alle Fische sind gleich schnell, nämlich genau einen Pixel mit jedem Bild. Und
das Allerschlimmste ist, das Ergebnis sieht auf jedem Computer anders aus, denn
die sogenannte Frame Rate (also wieviele Bilder pro Sekunde dargestellt werden)
hängt von der Leistungsfähigkeit des Rechners ab.

Mache die Fische unterschiedlich schnell und unabhängig von der Frame Rate.

Das Problem läßt sich ähnlich lösen wie die Geschwindigkeit der Animationen.

Tipp:
Da es ein sehr allgemeines Problem ist, das sowohl für die Animation als auch
für die Bewegung der Figuren und evt. später noch anderswo gelöst werden muss,
und es auch schon in zwei Typen `ImmerGeradeaus` und `HinUndHer` auftritt,
lohnt es sich darüber nachzudenken, ob und wie die Lösung sich in einem
separaten Typ kapseln läßt (z.B. namens `Geschwindigkeit`).

