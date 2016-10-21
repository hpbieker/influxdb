/*

Series:

	╔══════Series List═════╗
	║ ┌───────────────────┐║
	║ │     Term List     │║
	║ ├───────────────────┤║
	║ │    Series Data    │║
	║ ├───────────────────┤║
	║ │      Trailer      │║
	║ └───────────────────┘║
	╚══════════════════════╝

	╔══════════Term List═══════════╗
	║ ┌──────────────────────────┐ ║
	║ │   Term Count <uint32>    │ ║
	║ └──────────────────────────┘ ║
	║ ┏━━━━━━━━━━━━━━━━━━━━━━━━━━┓ ║
	║ ┃ ┌──────────────────────┐ ┃ ║
	║ ┃ │  len(Term) <varint>  │ ┃ ║
	║ ┃ ├──────────────────────┤ ┃ ║
	║ ┃ │    Term <byte...>    │ ┃ ║
	║ ┃ └──────────────────────┘ ┃ ║
	║ ┗━━━━━━━━━━━━━━━━━━━━━━━━━━┛ ║
	║ ┏━━━━━━━━━━━━━━━━━━━━━━━━━━┓ ║
	║ ┃ ┌──────────────────────┐ ┃ ║
	║ ┃ │  len(Term) <varint>  │ ┃ ║
	║ ┃ ├──────────────────────┤ ┃ ║
	║ ┃ │    Term <byte...>    │ ┃ ║
	║ ┃ └──────────────────────┘ ┃ ║
	║ ┗━━━━━━━━━━━━━━━━━━━━━━━━━━┛ ║
	╚══════════════════════════════╝

	╔═════════Series Data══════════╗
	║ ┌──────────────────────────┐ ║
	║ │  Series Count <uint32>   │ ║
	║ └──────────────────────────┘ ║
	║ ┏━━━━━━━━━━━━━━━━━━━━━━━━━━┓ ║
	║ ┃ ┌──────────────────────┐ ┃ ║
	║ ┃ │     Flag <uint8>     │ ┃ ║
	║ ┃ ├──────────────────────┤ ┃ ║
	║ ┃ │ len(Series) <varint> │ ┃ ║
	║ ┃ ├──────────────────────┤ ┃ ║
	║ ┃ │   Series <byte...>   │ ┃ ║
	║ ┃ └──────────────────────┘ ┃ ║
	║ ┗━━━━━━━━━━━━━━━━━━━━━━━━━━┛ ║
	║             ...              ║
	╚══════════════════════════════╝

	╔════════════Trailer══════════════╗
	║ ┌─────────────────────────────┐ ║
	║ │  Term List Offset <uint64>  │ ║
	║ ├─────────────────────────────┤ ║
	║ │   Term List Size <uint64>   │ ║
	║ ├─────────────────────────────┤ ║
	║ │ Series Data Offset <uint64> │ ║
	║ ├─────────────────────────────┤ ║
	║ │  Series Data Pos <uint64>   │ ║
	║ └─────────────────────────────┘ ║
	╚═════════════════════════════════╝


Tag Sets:

	╔════════Tag Set═════════╗
	║┌──────────────────────┐║
	║│   Tag Values Block   │║
	║├──────────────────────┤║
	║│         ...          │║
	║├──────────────────────┤║
	║│    Tag Keys Block    │║
	║├──────────────────────┤║
	║│       Trailer        │║
	║└──────────────────────┘║
	╚════════════════════════╝

	╔═══════Tag Values Block═══════╗
	║                              ║
	║ ┏━━━━━━━━Value List━━━━━━━━┓ ║
	║ ┃                          ┃ ║
	║ ┃┏━━━━━━━━━Value━━━━━━━━━━┓┃ ║
	║ ┃┃┌──────────────────────┐┃┃ ║
	║ ┃┃│     Flag <uint8>     │┃┃ ║
	║ ┃┃├──────────────────────┤┃┃ ║
	║ ┃┃│ len(Value) <varint>  │┃┃ ║
	║ ┃┃├──────────────────────┤┃┃ ║
	║ ┃┃│   Value <byte...>    │┃┃ ║
	║ ┃┃├──────────────────────┤┃┃ ║
	║ ┃┃│ len(Series) <varint> │┃┃ ║
	║ ┃┃├──────────────────────┤┃┃ ║
	║ ┃┃│SeriesIDs <uint32...> │┃┃ ║
	║ ┃┃└──────────────────────┘┃┃ ║
	║ ┃┗━━━━━━━━━━━━━━━━━━━━━━━━┛┃ ║
	║ ┃           ...            ┃ ║
	║ ┗━━━━━━━━━━━━━━━━━━━━━━━━━━┛ ║
	║ ┏━━━━━━━━Hash Index━━━━━━━━┓ ║
	║ ┃ ┌──────────────────────┐ ┃ ║
	║ ┃ │ len(Values) <uint32> │ ┃ ║
	║ ┃ ├──────────────────────┤ ┃ ║
	║ ┃ │Value Offset <uint64> │ ┃ ║
	║ ┃ ├──────────────────────┤ ┃ ║
	║ ┃ │         ...          │ ┃ ║
	║ ┃ └──────────────────────┘ ┃ ║
	║ ┗━━━━━━━━━━━━━━━━━━━━━━━━━━┛ ║
	╚══════════════════════════════╝

	╔════════Tag Key Block═════════╗
	║                              ║
	║ ┏━━━━━━━━━Key List━━━━━━━━━┓ ║
	║ ┃                          ┃ ║
	║ ┃┏━━━━━━━━━━Key━━━━━━━━━━━┓┃ ║
	║ ┃┃┌──────────────────────┐┃┃ ║
	║ ┃┃│     Flag <uint8>     │┃┃ ║
	║ ┃┃├──────────────────────┤┃┃ ║
	║ ┃┃│Value Offset <uint64> │┃┃ ║
	║ ┃┃├──────────────────────┤┃┃ ║
	║ ┃┃│  len(Key) <varint>   │┃┃ ║
	║ ┃┃├──────────────────────┤┃┃ ║
	║ ┃┃│    Key <byte...>     │┃┃ ║
	║ ┃┃└──────────────────────┘┃┃ ║
	║ ┃┗━━━━━━━━━━━━━━━━━━━━━━━━┛┃ ║
	║ ┃           ...            ┃ ║
	║ ┗━━━━━━━━━━━━━━━━━━━━━━━━━━┛ ║
	║ ┏━━━━━━━━Hash Index━━━━━━━━┓ ║
	║ ┃ ┌──────────────────────┐ ┃ ║
	║ ┃ │  len(Keys) <uint32>  │ ┃ ║
	║ ┃ ├──────────────────────┤ ┃ ║
	║ ┃ │ Key Offset <uint64>  │ ┃ ║
	║ ┃ ├──────────────────────┤ ┃ ║
	║ ┃ │         ...          │ ┃ ║
	║ ┃ └──────────────────────┘ ┃ ║
	║ ┗━━━━━━━━━━━━━━━━━━━━━━━━━━┛ ║
	╚══════════════════════════════╝

	╔════════════Trailer══════════════╗
	║ ┌─────────────────────────────┐ ║
	║ │  Hash Index Offset <uint64> │ ║
	║ ├─────────────────────────────┤ ║
	║ │    Tag Set Size <uint64>    │ ║
	║ ├─────────────────────────────┤ ║
	║ │   Tag Set Version <uint16>  │ ║
	║ └─────────────────────────────┘ ║
	╚═════════════════════════════════╝


Measurements

	╔══════════Measurements Block═══════════╗
	║                                       ║
	║ ┏━━━━━━━━━Measurement List━━━━━━━━━━┓ ║
	║ ┃                                   ┃ ║
	║ ┃┏━━━━━━━━━━Measurement━━━━━━━━━━━┓ ┃ ║
	║ ┃┃┌─────────────────────────────┐ ┃ ┃ ║
	║ ┃┃│        Flag <uint8>         │ ┃ ┃ ║
	║ ┃┃├─────────────────────────────┤ ┃ ┃ ║
	║ ┃┃│  Tag Block Offset <uint64>  │ ┃ ┃ ║
	║ ┃┃├─────────────────────────────┤ ┃ ┃ ║
	║ ┃┃│     len(Name) <varint>      │ ┃ ┃ ║
	║ ┃┃├─────────────────────────────┤ ┃ ┃ ║
	║ ┃┃│       Name <byte...>        │ ┃ ┃ ║
	║ ┃┃├─────────────────────────────┤ ┃ ┃ ║
	║ ┃┃│    len(Series) <uint32>     │ ┃ ┃ ║
	║ ┃┃├─────────────────────────────┤ ┃ ┃ ║
	║ ┃┃│    SeriesIDs <uint32...>    │ ┃ ┃ ║
	║ ┃┃└─────────────────────────────┘ ┃ ┃ ║
	║ ┃┗━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┛ ┃ ║
	║ ┃                ...                ┃ ║
	║ ┗━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┛ ║
	║ ┏━━━━━━━━━━━━Hash Index━━━━━━━━━━━━━┓ ║
	║ ┃ ┌───────────────────────────────┐ ┃ ║
	║ ┃ │  len(Measurements) <uint32>   │ ┃ ║
	║ ┃ ├───────────────────────────────┤ ┃ ║
	║ ┃ │  Measurement Offset <uint64>  │ ┃ ║
	║ ┃ ├───────────────────────────────┤ ┃ ║
	║ ┃ │              ...              │ ┃ ║
	║ ┃ └───────────────────────────────┘ ┃ ║
	║ ┗━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┛ ║
	║ ┏━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┓ ║
	║ ┃              Trailer              ┃ ║
	║ ┗━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┛ ║
	╚═══════════════════════════════════════╝

	╔════════════Trailer══════════════╗
	║ ┌─────────────────────────────┐ ║
	║ │  Hash Index Offset <uint64> │ ║
	║ ├─────────────────────────────┤ ║
	║ │     Block Size <uint64>     │ ║
	║ ├─────────────────────────────┤ ║
	║ │    Block Version <uint16>   │ ║
	║ └─────────────────────────────┘ ║
	╚═════════════════════════════════╝


WAL

	╔═════════════WAL══════════════╗
	║                              ║
	║ ┏━━━━━━━━━━Entry━━━━━━━━━━━┓ ║
	║ ┃ ┌──────────────────────┐ ┃ ║
	║ ┃ │     Flag <uint8>     │ ┃ ║
	║ ┃ ├──────────────────────┤ ┃ ║
	║ ┃ │  len(Name) <varint>  │ ┃ ║
	║ ┃ ├──────────────────────┤ ┃ ║
	║ ┃ │    Name <byte...>    │ ┃ ║
	║ ┃ ├──────────────────────┤ ┃ ║
	║ ┃ │  len(Tags) <varint>  │ ┃ ║
	║ ┃ ├──────────────────────┤ ┃ ║
	║ ┃ │  len(Key0) <varint>  │ ┃ ║
	║ ┃ ├──────────────────────┤ ┃ ║
	║ ┃ │    Key0 <byte...>    │ ┃ ║
	║ ┃ ├──────────────────────┤ ┃ ║
	║ ┃ │ len(Value0) <varint> │ ┃ ║
	║ ┃ ├──────────────────────┤ ┃ ║
	║ ┃ │   Value0 <byte...>   │ ┃ ║
	║ ┃ ├──────────────────────┤ ┃ ║
	║ ┃ │         ...          │ ┃ ║
	║ ┃ ├──────────────────────┤ ┃ ║
	║ ┃ │  Checksum <uint32>   │ ┃ ║
	║ ┃ └──────────────────────┘ ┃ ║
	║ ┗━━━━━━━━━━━━━━━━━━━━━━━━━━┛ ║
	║             ...              ║
	╚══════════════════════════════╝


*/
package tsi1