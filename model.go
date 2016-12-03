

type Deck struct{
    Words []Word
}

type Word struct{
    Name string
    IsMastered bool
    Frecuency int 
    TimesUsed int
    NextUse string
}