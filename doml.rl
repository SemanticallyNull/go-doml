package doml

import "fmt"

%%{
	machine doml;
	write data;
}%%

func Parse(data string) (map[string]any, error) {
  cs, p, pe := 0, 0, len(data)
  val := make(map[string]any, 0)
  var key, current string
  currentArr := make([]string, 0)

  %%{
    whitespace = [\n\t ];
    noncomma = [^◀️];
    normal = noncomma & [^✌️🛑▶️\n\t ];

    action separator { val = append(val, current); current = "" }
    action character { current = current + string(fc) }
    action endkey { key = current; current = "" }
    action endsinglevalue { val[key] = current; current = "" }
    action endarraysinglevalue { currentArr = append(currentArr, current); current = "" }
    action endarrayvalue { val[key] = currentArr; currentArr = make([]string, 0) }

    single = '▶️' ( normal @character )* '◀️';
    array = '⏩' whitespace* ( single @endarraysinglevalue whitespace* )* '⏪';
    key = '✌️' ( normal @character )* '✌️';
    kvpair = ( key @endkey ) whitespace* (( single @endsinglevalue) | ( array @endarrayvalue )) whitespace*;
    main := whitespace* '🏁' whitespace* ( kvpair )* whitespace* '🛑' whitespace*;

    write init;
    write exec;
  }%%

  if cs < doml_first_final {
    return map[string]any{}, fmt.Errorf("parse error: %s", data)
  }

  return val, nil
}
