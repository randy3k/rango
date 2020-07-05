#ifndef R_H__
#define R_H__

# include <stddef.h>

#ifndef RANGO
#define EXTERN extern
#else
#define EXTERN
#endif

typedef unsigned char Rbyte;

typedef enum {Bytes, Chars, Width} nchar_type;
typedef void * (*DL_FUNC)();

typedef int R_len_t;

// FIXME: when size of size_t <= 4
typedef ptrdiff_t R_xlen_t;

typedef unsigned int SEXPTYPE;

static const unsigned int NILSXP     =  0;
static const unsigned int SYMSXP     =  1;
static const unsigned int LISTSXP    =  2;
static const unsigned int CLOSXP     =  3;
static const unsigned int ENVSXP     =  4;
static const unsigned int PROMSXP    =  5;
static const unsigned int LANGSXP    =  6;
static const unsigned int SPECIALSXP =  7;
static const unsigned int BUILTINSXP =  8;
static const unsigned int CHARSXP    =  9;
static const unsigned int LGLSXP     = 10;
static const unsigned int INTSXP     = 13;
static const unsigned int REALSXP    = 14;
static const unsigned int CPLXSXP    = 15;
static const unsigned int STRSXP     = 16;
static const unsigned int DOTSXP     = 17;
static const unsigned int ANYSXP     = 18;
static const unsigned int VECSXP     = 19;
static const unsigned int EXPRSXP    = 20;
static const unsigned int BCODESXP   = 21;
static const unsigned int EXTPTRSXP  = 22;
static const unsigned int WEAKREFSXP = 23;
static const unsigned int RAWSXP     = 24;
static const unsigned int S4SXP      = 25;
static const unsigned int NEWSXP     = 30;
static const unsigned int FREESXP    = 31;
static const unsigned int FUNSXP     = 99;

typedef struct {
    double r;
    double i;
} Rcomplex;

typedef enum { FALSE = 0, TRUE } Rboolean;

struct sxpinfo_struct {
    SEXPTYPE type      :  5;
    unsigned int scalar:  1;
    unsigned int obj   :  1;
    unsigned int alt   :  1;
    unsigned int gp    : 16;
    unsigned int mark  :  1;
    unsigned int debug :  1;
    unsigned int trace :  1;
    unsigned int spare :  1;
    unsigned int gcgen :  1;
    unsigned int gccls :  3;
    unsigned int named : 16;
    unsigned int extra : 16;
};

struct vecsxp_struct {
    R_xlen_t    length;
    R_xlen_t    truelength;
};


struct primsxp_struct {
    int offset;
};

struct symsxp_struct {
    struct SEXPREC *pname;
    struct SEXPREC *value;
    struct SEXPREC *internal;
};

struct listsxp_struct {
    struct SEXPREC *carval;
    struct SEXPREC *cdrval;
    struct SEXPREC *tagval;
};

struct envsxp_struct {
    struct SEXPREC *frame;
    struct SEXPREC *enclos;
    struct SEXPREC *hashtab;
};

struct closxp_struct {
    struct SEXPREC *formals;
    struct SEXPREC *body;
    struct SEXPREC *env;
};

struct promsxp_struct {
    struct SEXPREC *value;
    struct SEXPREC *expr;
    struct SEXPREC *env;
};

typedef struct SEXPREC {
    struct sxpinfo_struct sxpinfo;
    struct SEXPREC *attrib;
    struct SEXPREC *gengc_next_node, *gengc_prev_node;
    union {
        struct primsxp_struct primsxp;
        struct symsxp_struct symsxp;
        struct listsxp_struct listsxp;
        struct envsxp_struct envsxp;
        struct closxp_struct closxp;
        struct promsxp_struct promsxp;
    } u;
} SEXPREC;

typedef struct SEXPREC *SEXP;

// Rinternals.h
EXTERN const char* (*R_CHAR)(SEXP x);
EXTERN Rboolean (*Rf_isNull)(SEXP s);
EXTERN Rboolean (*Rf_isSymbol)(SEXP s);
EXTERN Rboolean (*Rf_isLogical)(SEXP s);
EXTERN Rboolean (*Rf_isReal)(SEXP s);
EXTERN Rboolean (*Rf_isComplex)(SEXP s);
EXTERN Rboolean (*Rf_isExpression)(SEXP s);
EXTERN Rboolean (*Rf_isEnvironment)(SEXP s);
EXTERN Rboolean (*Rf_isString)(SEXP s);
EXTERN Rboolean (*Rf_isObject)(SEXP s);

EXTERN int (*TYPEOF)(SEXP x);
EXTERN int (*IS_S4_OBJECT)(SEXP x);

EXTERN int  (*LENGTH)(SEXP x);
EXTERN R_xlen_t (*XLENGTH)(SEXP x);
EXTERN R_xlen_t  (*TRUELENGTH)(SEXP x);
EXTERN void (*SETLENGTH)(SEXP x, R_xlen_t v);
EXTERN void (*SET_TRUELENGTH)(SEXP x, R_xlen_t v);
EXTERN int  (*IS_LONG_VEC)(SEXP x);
EXTERN int  (*LEVELS)(SEXP x);
EXTERN int  (*SETLEVELS)(SEXP x, int v);

// Vector Access Functions
EXTERN int *(*LOGICAL)(SEXP x);
EXTERN int  *(*INTEGER)(SEXP x);
EXTERN Rbyte *(*RAW)(SEXP x);
EXTERN double *(*REAL)(SEXP x);
EXTERN Rcomplex *(*COMPLEX)(SEXP x);
EXTERN SEXP (*STRING_ELT)(SEXP x, R_xlen_t i);
EXTERN SEXP (*VECTOR_ELT)(SEXP x, R_xlen_t i);
EXTERN void (*SET_STRING_ELT)(SEXP x, R_xlen_t i, SEXP v);
EXTERN SEXP (*SET_VECTOR_ELT)(SEXP x, R_xlen_t i, SEXP v);

// List Access
EXTERN SEXP (*Rf_cons)(SEXP, SEXP);
EXTERN SEXP (*Rf_lcons)(SEXP, SEXP);
EXTERN SEXP (*TAG)(SEXP e);
EXTERN SEXP (*CAR)(SEXP e);
EXTERN SEXP (*CDR)(SEXP e);
EXTERN SEXP (*CAAR)(SEXP e);
EXTERN SEXP (*CDAR)(SEXP e);
EXTERN SEXP (*CADR)(SEXP e);
EXTERN SEXP (*CDDR)(SEXP e);
EXTERN SEXP (*CDDDR)(SEXP e);
EXTERN SEXP (*CADDR)(SEXP e);
EXTERN SEXP (*CADDDR)(SEXP e);
EXTERN SEXP (*CAD4R)(SEXP e);
EXTERN int  (*MISSING)(SEXP x);
EXTERN void (*SET_MISSING)(SEXP x, int v);
EXTERN void (*SET_TAG)(SEXP x, SEXP y);
EXTERN SEXP (*SETCAR)(SEXP x, SEXP y);
EXTERN SEXP (*SETCDR)(SEXP x, SEXP y);
EXTERN SEXP (*SETCADR)(SEXP x, SEXP y);
EXTERN SEXP (*SETCADDR)(SEXP x, SEXP y);
EXTERN SEXP (*SETCADDDR)(SEXP x, SEXP y);
EXTERN SEXP (*SETCAD4R)(SEXP e, SEXP y);
EXTERN SEXP (*CONS_NR)(SEXP a, SEXP b);

EXTERN SEXP (*PRINTNAME)(SEXP x);

EXTERN SEXP (*Rf_protect)(SEXP);
EXTERN void (*Rf_unprotect)(int);

EXTERN SEXP R_GlobalEnv;
EXTERN SEXP R_EmptyEnv;
EXTERN SEXP R_BaseEnv;
EXTERN SEXP R_BaseNamespace;
EXTERN SEXP R_NamespaceRegistry;
EXTERN SEXP R_Srcref;
EXTERN SEXP R_NilValue;
EXTERN SEXP R_UnboundValue;
EXTERN SEXP R_MissingArg;
EXTERN SEXP R_InBCInterpreter;
EXTERN SEXP R_CurrentExpression;
EXTERN SEXP R_AsCharacterSymbol;
EXTERN SEXP R_baseSymbol;
EXTERN SEXP R_BaseSymbol;
EXTERN SEXP R_BraceSymbol;
EXTERN SEXP R_Bracket2Symbol;
EXTERN SEXP R_BracketSymbol;
EXTERN SEXP R_ClassSymbol;
EXTERN SEXP R_DeviceSymbol;
EXTERN SEXP R_DimNamesSymbol;
EXTERN SEXP R_DimSymbol;
EXTERN SEXP R_DollarSymbol;
EXTERN SEXP R_DotsSymbol;
EXTERN SEXP R_DoubleColonSymbol;
EXTERN SEXP R_DropSymbol;
EXTERN SEXP R_LastvalueSymbol;
EXTERN SEXP R_LevelsSymbol;
EXTERN SEXP R_ModeSymbol;
EXTERN SEXP R_NaRmSymbol;
EXTERN SEXP R_NameSymbol;
EXTERN SEXP R_NamesSymbol;
EXTERN SEXP R_NamespaceEnvSymbol;
EXTERN SEXP R_PackageSymbol;
EXTERN SEXP R_PreviousSymbol;
EXTERN SEXP R_QuoteSymbol;
EXTERN SEXP R_RowNamesSymbol;
EXTERN SEXP R_SeedsSymbol;
EXTERN SEXP R_SortListSymbol;
EXTERN SEXP R_SourceSymbol;
EXTERN SEXP R_SpecSymbol;
EXTERN SEXP R_TripleColonSymbol;
EXTERN SEXP R_TspSymbol;
EXTERN SEXP R_dot_defined;
EXTERN SEXP R_dot_Method;
EXTERN SEXP R_dot_packageName;
EXTERN SEXP R_dot_target;
EXTERN SEXP R_dot_Generic;
EXTERN SEXP R_NaString;
EXTERN SEXP R_BlankString;
EXTERN SEXP R_BlankScalarString;

EXTERN SEXP (*Rf_asChar)(SEXP);
EXTERN SEXP (*Rf_coerceVector)(SEXP, SEXPTYPE);
EXTERN SEXP (*Rf_PairToVectorList)(SEXP x);
EXTERN SEXP (*Rf_VectorToPairList)(SEXP x);
EXTERN SEXP (*Rf_asCharacterFactor)(SEXP x);
EXTERN int (*Rf_asLogical)(SEXP x);
EXTERN int (*Rf_asLogical2)(SEXP x, int checking, SEXP call, SEXP rho);
EXTERN int (*Rf_asInteger)(SEXP x);
EXTERN double (*Rf_asReal)(SEXP x);
EXTERN Rcomplex (*Rf_asComplex)(SEXP x);

EXTERN char * (*Rf_acopy_string)(const char *);
// EXTERN void (*Rf_addMissingVarsToNewEnv)(SEXP, SEXP);
EXTERN SEXP (*Rf_alloc3DArray)(SEXPTYPE, int, int, int);
EXTERN SEXP (*Rf_allocArray)(SEXPTYPE, SEXP);
EXTERN SEXP (*Rf_allocFormalsList2)(SEXP sym1, SEXP sym2);
EXTERN SEXP (*Rf_allocFormalsList3)(SEXP sym1, SEXP sym2, SEXP sym3);
EXTERN SEXP (*Rf_allocFormalsList4)(SEXP sym1, SEXP sym2, SEXP sym3, SEXP sym4);
EXTERN SEXP (*Rf_allocFormalsList5)(SEXP sym1, SEXP sym2, SEXP sym3, SEXP sym4, SEXP sym5);
EXTERN SEXP (*Rf_allocFormalsList6)(SEXP sym1, SEXP sym2, SEXP sym3, SEXP sym4, SEXP sym5, SEXP sym6);
EXTERN SEXP (*Rf_allocMatrix)(SEXPTYPE, int, int);
EXTERN SEXP (*Rf_allocList)(int);
EXTERN SEXP (*Rf_allocS4Object)(void);
EXTERN SEXP (*Rf_allocSExp)(SEXPTYPE);
EXTERN SEXP (*Rf_allocVector3)(SEXPTYPE, R_xlen_t, void*);
EXTERN R_xlen_t (*Rf_any_duplicated)(SEXP x, Rboolean from_last);
EXTERN R_xlen_t (*Rf_any_duplicated3)(SEXP x, SEXP incomp, Rboolean from_last);
EXTERN SEXP (*Rf_applyClosure)(SEXP, SEXP, SEXP, SEXP, SEXP);
EXTERN SEXP (*Rf_arraySubscript)(int, SEXP, SEXP, SEXP (*)(SEXP,SEXP), SEXP (*)(SEXP, int), SEXP);
EXTERN SEXP (*Rf_classgets)(SEXP, SEXP);
EXTERN void (*Rf_copyMatrix)(SEXP, SEXP, Rboolean);
EXTERN void (*Rf_copyListMatrix)(SEXP, SEXP, Rboolean);
EXTERN void (*Rf_copyMostAttrib)(SEXP, SEXP);
EXTERN void (*Rf_copyVector)(SEXP, SEXP);
EXTERN int (*Rf_countContexts)(int, int);
EXTERN SEXP (*Rf_CreateTag)(SEXP);
EXTERN void (*Rf_defineVar)(SEXP, SEXP, SEXP);
EXTERN SEXP (*Rf_dimgets)(SEXP, SEXP);
EXTERN SEXP (*Rf_dimnamesgets)(SEXP, SEXP);
EXTERN SEXP (*Rf_DropDims)(SEXP);
EXTERN SEXP (*Rf_duplicate)(SEXP);
EXTERN SEXP (*Rf_shallow_duplicate)(SEXP);
// EXTERN SEXP (*R_duplicate_attr)(SEXP);
// EXTERN SEXP (*R_shallow_duplicate_attr)(SEXP);
EXTERN SEXP (*Rf_lazy_duplicate)(SEXP);

EXTERN SEXP (*Rf_duplicated)(SEXP, Rboolean);
EXTERN Rboolean (*R_envHasNoSpecialSymbols)(SEXP);
EXTERN SEXP (*Rf_eval)(SEXP, SEXP);
EXTERN SEXP (*Rf_findFun)(SEXP, SEXP);
EXTERN SEXP (*Rf_findVar)(SEXP, SEXP);
EXTERN SEXP (*Rf_findVarInFrame)(SEXP, SEXP);
EXTERN SEXP (*Rf_findVarInFrame3)(SEXP, SEXP, Rboolean);
EXTERN SEXP (*Rf_getAttrib)(SEXP, SEXP);
EXTERN SEXP (*Rf_GetArrayDimnames)(SEXP);
EXTERN SEXP (*Rf_GetColNames)(SEXP);
EXTERN void (*Rf_GetMatrixDimnames)(SEXP, SEXP*, SEXP*, const char**, const char**);
EXTERN SEXP (*Rf_GetOption1)(SEXP);
EXTERN int (*Rf_GetOptionDigits)(void);
EXTERN int (*Rf_GetOptionWidth)(void);
EXTERN SEXP (*Rf_GetRowNames)(SEXP);
EXTERN void (*Rf_gsetVar)(SEXP, SEXP, SEXP);
EXTERN SEXP (*Rf_install)(const char *);
EXTERN SEXP (*Rf_installChar)(SEXP);
EXTERN Rboolean (*Rf_isFree)(SEXP);
EXTERN Rboolean (*Rf_isOrdered)(SEXP);
EXTERN Rboolean (*Rf_isUnordered)(SEXP);
EXTERN Rboolean (*Rf_isUnsorted)(SEXP, Rboolean);
EXTERN SEXP (*Rf_lengthgets)(SEXP, R_len_t);
EXTERN SEXP (*Rf_xlengthgets)(SEXP, R_xlen_t);
EXTERN SEXP (*R_lsInternal)(SEXP, Rboolean);
EXTERN SEXP (*R_lsInternal3)(SEXP, Rboolean, Rboolean);
EXTERN SEXP (*Rf_match)(SEXP, SEXP, int);
EXTERN SEXP (*Rf_matchE)(SEXP, SEXP, int, SEXP);
EXTERN SEXP (*Rf_namesgets)(SEXP, SEXP);
EXTERN SEXP (*Rf_mkChar)(const char *);
EXTERN SEXP (*Rf_mkCharLen)(const char *, int);
EXTERN Rboolean (*Rf_NonNullStringMatch)(SEXP, SEXP);
EXTERN int (*Rf_ncols)(SEXP);
EXTERN int (*Rf_nrows)(SEXP);
EXTERN SEXP (*Rf_nthcdr)(SEXP, int);

EXTERN int (*R_nchar)(SEXP string, nchar_type type_, Rboolean allowNA, Rboolean keepNA, const char* msg_name);
EXTERN Rboolean (*Rf_pmatch)(SEXP, SEXP, Rboolean);
EXTERN Rboolean (*Rf_psmatch)(const char *, const char *, Rboolean);
EXTERN SEXP (*R_ParseEvalString)(const char *, SEXP);
EXTERN void (*Rf_PrintValue)(SEXP);
// EXTERN void (*Rf_printwhere)(void);
// EXTERN void (*Rf_readS3VarsFromFrame)(SEXP, SEXP*, SEXP*, SEXP*, SEXP*, SEXP*, SEXP*);
EXTERN SEXP (*Rf_setAttrib)(SEXP, SEXP, SEXP);
EXTERN void (*Rf_setSVector)(SEXP*, int, SEXP);
EXTERN void (*Rf_setVar)(SEXP, SEXP, SEXP);
// EXTERN SEXP (*Rf_stringSuffix)(SEXP, int);
EXTERN SEXPTYPE (*Rf_str2type)(const char *);
EXTERN Rboolean (*Rf_StringBlank)(SEXP);
EXTERN SEXP (*Rf_substitute)(SEXP,SEXP);
EXTERN SEXP (*Rf_topenv)(SEXP, SEXP);
EXTERN const char * (*Rf_translateChar)(SEXP);
EXTERN const char * (*Rf_translateChar0)(SEXP);
EXTERN const char * (*Rf_translateCharUTF8)(SEXP);
EXTERN const char * (*Rf_type2char)(SEXPTYPE);
EXTERN SEXP (*Rf_type2rstr)(SEXPTYPE);
EXTERN SEXP (*Rf_type2str)(SEXPTYPE);
EXTERN SEXP (*Rf_type2str_nowarn)(SEXPTYPE);

EXTERN SEXP (*R_tryEval)(SEXP, SEXP, int *);
EXTERN SEXP (*R_tryEvalSilent)(SEXP, SEXP, int *);
EXTERN const char *(*R_curErrorBuf)();

EXTERN Rboolean (*Rf_isS4)(SEXP);
EXTERN SEXP (*Rf_asS4)(SEXP, Rboolean, int);
EXTERN SEXP (*Rf_S3Class)(SEXP);
EXTERN int (*Rf_isBasicClass)(const char *);

typedef enum {
    CE_NATIVE = 0,
    CE_UTF8   = 1,
    CE_LATIN1 = 2,
    CE_BYTES  = 3,
    CE_SYMBOL = 5,
    CE_ANY    =99
} cetype_t;

EXTERN cetype_t (*Rf_getCharCE)(SEXP);
EXTERN SEXP (*Rf_mkCharCE)(const char *, cetype_t);
EXTERN SEXP (*Rf_mkCharLenCE)(const char *, int, cetype_t);
EXTERN const char *(*Rf_reEnc)(const char *x, cetype_t ce_in, cetype_t ce_out, int subst);

EXTERN SEXP (*R_MakeExternalPtr)(void *p, SEXP tag, SEXP prot);
EXTERN void *(*R_ExternalPtrAddr)(SEXP s);
EXTERN SEXP (*R_ExternalPtrTag)(SEXP s);
EXTERN SEXP (*R_ExternalPtrProtected)(SEXP s);
EXTERN void (*R_ClearExternalPtr)(SEXP s);
EXTERN void (*R_SetExternalPtrAddr)(SEXP s, void *p);
EXTERN void (*R_SetExternalPtrTag)(SEXP s, SEXP tag);
EXTERN void (*R_SetExternalPtrProtected)(SEXP s, SEXP p);
EXTERN SEXP (*R_MakeExternalPtrFn)(DL_FUNC p, SEXP tag, SEXP prot);
EXTERN DL_FUNC (*R_ExternalPtrAddrFn)(SEXP s);

typedef void (*R_CFinalizer_t)(SEXP);

EXTERN void (*R_RegisterFinalizer)(SEXP s, SEXP fun);
EXTERN void (*R_RegisterCFinalizer)(SEXP s, R_CFinalizer_t fun);
EXTERN void (*R_RegisterFinalizerEx)(SEXP s, SEXP fun, Rboolean onexit);
EXTERN void (*R_RegisterCFinalizerEx)(SEXP s, R_CFinalizer_t fun, Rboolean onexit);
EXTERN void (*R_RunPendingFinalizers)(void);

EXTERN Rboolean (*R_ToplevelExec)(void (*fun)(void *), void *data);
EXTERN SEXP (*R_tryCatch)(SEXP (*)(void *), void *, SEXP, SEXP (*)(SEXP, void *), void *, void (*)(void *), void *);
EXTERN SEXP (*R_tryCatchError)(SEXP (*)(void *), void *, SEXP (*)(SEXP, void *), void *);

EXTERN void (*R_RestoreHashCount)(SEXP rho);
EXTERN Rboolean (*R_IsPackageEnv)(SEXP rho);
EXTERN SEXP (*R_PackageEnvName)(SEXP rho);
EXTERN SEXP (*R_FindPackageEnv)(SEXP info);
EXTERN Rboolean (*R_IsNamespaceEnv)(SEXP rho);
EXTERN SEXP (*R_NamespaceEnvSpec)(SEXP rho);
EXTERN SEXP (*R_FindNamespace)(SEXP info);
EXTERN void (*R_LockEnvironment)(SEXP env, Rboolean bindings);
EXTERN Rboolean (*R_EnvironmentIsLocked)(SEXP env);
EXTERN void (*R_LockBinding)(SEXP sym, SEXP env);
EXTERN void (*R_unLockBinding)(SEXP sym, SEXP env);
EXTERN void (*R_MakeActiveBinding)(SEXP sym, SEXP fun, SEXP env);
EXTERN Rboolean (*R_BindingIsLocked)(SEXP sym, SEXP env);
EXTERN Rboolean (*R_BindingIsActive)(SEXP sym, SEXP env);
EXTERN Rboolean (*R_HasFancyBindings)(SEXP rho);

EXTERN void (*Rf_errorcall)(SEXP, const char *, ...);
EXTERN void (*Rf_warningcall)(SEXP, const char *, ...);

EXTERN SEXP (*R_do_slot)(SEXP obj, SEXP name);
EXTERN SEXP (*R_do_slot_assign)(SEXP obj, SEXP name, SEXP value);
EXTERN int (*R_has_slot)(SEXP obj, SEXP name);
EXTERN SEXP (*R_S4_extends)(SEXP klass, SEXP useTable);

EXTERN void (*R_PreserveObject)(SEXP);
EXTERN void (*R_ReleaseObject)(SEXP);

EXTERN void (*R_dot_Last)(void);
EXTERN void (*R_RunExitFinalizers)(void);

EXTERN Rboolean (*R_compute_identical)(SEXP, SEXP, int);

EXTERN SEXP     (*Rf_allocVector)(SEXPTYPE, R_xlen_t);
EXTERN Rboolean (*Rf_conformable)(SEXP, SEXP);
EXTERN SEXP     (*Rf_elt)(SEXP, int);
EXTERN Rboolean (*Rf_inherits)(SEXP, const char *);
EXTERN Rboolean (*Rf_isArray)(SEXP);
EXTERN Rboolean (*Rf_isFactor)(SEXP);
EXTERN Rboolean (*Rf_isFrame)(SEXP);
EXTERN Rboolean (*Rf_isFunction)(SEXP);
EXTERN Rboolean (*Rf_isInteger)(SEXP);
EXTERN Rboolean (*Rf_isLanguage)(SEXP);
EXTERN Rboolean (*Rf_isList)(SEXP);
EXTERN Rboolean (*Rf_isMatrix)(SEXP);
EXTERN Rboolean (*Rf_isNewList)(SEXP);
EXTERN Rboolean (*Rf_isNumber)(SEXP);
EXTERN Rboolean (*Rf_isNumeric)(SEXP);
EXTERN Rboolean (*Rf_isPairList)(SEXP);
EXTERN Rboolean (*Rf_isPrimitive)(SEXP);
EXTERN Rboolean (*Rf_isTs)(SEXP);
EXTERN Rboolean (*Rf_isUserBinop)(SEXP);
EXTERN Rboolean (*Rf_isValidString)(SEXP);
EXTERN Rboolean (*Rf_isValidStringF)(SEXP);
EXTERN Rboolean (*Rf_isVector)(SEXP);
EXTERN Rboolean (*Rf_isVectorAtomic)(SEXP);
EXTERN Rboolean (*Rf_isVectorList)(SEXP);
EXTERN Rboolean (*Rf_isVectorizable)(SEXP);
EXTERN SEXP     (*Rf_lang1)(SEXP);
EXTERN SEXP     (*Rf_lang2)(SEXP, SEXP);
EXTERN SEXP     (*Rf_lang3)(SEXP, SEXP, SEXP);
EXTERN SEXP     (*Rf_lang4)(SEXP, SEXP, SEXP, SEXP);
EXTERN SEXP     (*Rf_lang5)(SEXP, SEXP, SEXP, SEXP, SEXP);
EXTERN SEXP     (*Rf_lang6)(SEXP, SEXP, SEXP, SEXP, SEXP, SEXP);
EXTERN SEXP     (*Rf_lastElt)(SEXP);
EXTERN R_len_t  (*Rf_length)(SEXP);
EXTERN SEXP     (*Rf_list1)(SEXP);
EXTERN SEXP     (*Rf_list2)(SEXP, SEXP);
EXTERN SEXP     (*Rf_list3)(SEXP, SEXP, SEXP);
EXTERN SEXP     (*Rf_list4)(SEXP, SEXP, SEXP, SEXP);
EXTERN SEXP     (*Rf_list5)(SEXP, SEXP, SEXP, SEXP, SEXP);
EXTERN SEXP     (*Rf_list6)(SEXP, SEXP, SEXP, SEXP, SEXP, SEXP);
EXTERN SEXP     (*Rf_listAppend)(SEXP, SEXP);
EXTERN SEXP     (*Rf_mkNamed)(SEXPTYPE, const char **);
EXTERN SEXP     (*Rf_mkString)(const char *);
EXTERN int  (*Rf_nlevels)(SEXP);
EXTERN int  (*Rf_stringPositionTr)(SEXP, const char *);
EXTERN SEXP     (*Rf_ScalarComplex)(Rcomplex);
EXTERN SEXP     (*Rf_ScalarInteger)(int);
EXTERN SEXP     (*Rf_ScalarLogical)(int);
EXTERN SEXP     (*Rf_ScalarRaw)(Rbyte);
EXTERN SEXP     (*Rf_ScalarReal)(double);
EXTERN SEXP     (*Rf_ScalarString)(SEXP);
EXTERN R_xlen_t  (*Rf_xlength)(SEXP);
EXTERN R_xlen_t  (*XTRUELENGTH)(SEXP x);
// EXTERN int (*LENGTH_EX)(SEXP x, const char *file, int line);
// EXTERN R_xlen_t (*XLENGTH_EX)(SEXP x);


// Arith.h
// EXTERN double R_NaN;
// EXTERN double R_PosInf;
// EXTERN double R_NegInf;
// EXTERN double R_NaReal;
// EXTERN int    R_NaInt;
EXTERN int (*R_IsNA)(double);
EXTERN int (*R_IsNaN)(double);
EXTERN int (*R_finite)(double);


// Parse.h
typedef enum {
    PARSE_NULL,
    PARSE_OK,
    PARSE_INCOMPLETE,
    PARSE_ERROR,
    PARSE_EOF
} ParseStatus;

EXTERN SEXP (*_R_ParseVector)(SEXP, int, ParseStatus *, SEXP);

// Memory.h
EXTERN void*   (*vmaxget)(void);
EXTERN void    (*vmaxset)(const void *);

EXTERN void    (*R_gc)(void);
EXTERN int (*R_gc_running)();

EXTERN char*   (*R_alloc)(size_t, int);
EXTERN long double *(*R_allocLD)(size_t nelem);

// EXTERN void *  (*R_malloc_gc)(size_t);
// EXTERN void *  (*R_calloc_gc)(size_t, size_t);
// EXTERN void *  (*R_realloc_gc)(void *, size_t);

// Error.h
EXTERN void    (*Rf_error)(const char *, ...);
EXTERN void    (*Rf_warning)(const char *, ...);
EXTERN void    (*R_ShowMessage)(const char *s);

// Defn.h
// EXTERN void (*Rf_CoercionWarning)(int);/* warning code */
// EXTERN int (*Rf_LogicalFromInteger)(int, int*);
// EXTERN int (*Rf_LogicalFromReal)(double, int*);
// EXTERN int (*Rf_LogicalFromComplex)(Rcomplex, int*);
// EXTERN int (*Rf_IntegerFromLogical)(int, int*);
// EXTERN int (*Rf_IntegerFromReal)(double, int*);
// EXTERN int (*Rf_IntegerFromComplex)(Rcomplex, int*);
// EXTERN double (*Rf_RealFromLogical)(int, int*);
// EXTERN double (*Rf_RealFromInteger)(int, int*);
// EXTERN double (*Rf_RealFromComplex)(Rcomplex, int*);
// EXTERN Rcomplex (*Rf_ComplexFromLogical)(int, int*);
// EXTERN Rcomplex (*Rf_ComplexFromInteger)(int, int*);
// EXTERN Rcomplex (*Rf_ComplexFromReal)(double, int*);

EXTERN void (*R_ProcessEvents)(void);

// EXTERN void (*Rf_PrintVersion)(char *, size_t len);
// EXTERN void (*Rf_PrintVersion_part_1)(char *, size_t len);
// EXTERN void (*Rf_PrintVersionString)(char *, size_t len);
EXTERN SEXP (*R_data_class)(SEXP , Rboolean);

// Utils.h
EXTERN void (*R_CheckUserInterrupt)(void);

// RStartup.h

typedef struct
{
    Rboolean R_Quiet;
    Rboolean R_Slave;
    Rboolean R_Interactive;
    Rboolean R_Verbose;
    Rboolean LoadSiteFile;
    Rboolean LoadInitFile;
    Rboolean DebugInitFile;
    int RestoreAction;
    int SaveAction;
    size_t vsize;
    size_t nsize;
    size_t max_vsize;
    size_t max_nsize;
    size_t ppsize;
    int NoRenviron;
    char *rhome;
    char *home;
    // we use _ReadConsole and _WriteConsole to avoid name collision
    int  (*_ReadConsole)(const char *, unsigned char *, int, int);
    void (*_WriteConsole)(const char *, int);
    void (*CallBack)(void);
    void (*ShowMessage) (const char *);
    int (*YesNoCancel) (const char *);
    void (*Busy) (int);
    int CharacterMode;
    void (*WriteConsoleEx)(const char *, int, int);
} structRstart;
typedef structRstart *Rstart;

EXTERN void (*R_DefParams)(Rstart);
EXTERN void (*R_SetParams)(Rstart);
EXTERN void (*R_set_command_line_arguments)(int argc, char **argv);

// Rinterface.h

EXTERN int (*Rstd_CleanUp)(int saveact, int status, int RunLast);
EXTERN int *R_SignalHandlers_t;


// Rembedded.h
EXTERN int (*Rf_initialize_R)(int ac, char **av);
EXTERN void (*setup_Rmainloop)(void);
EXTERN void (*run_Rmainloop)(void);

// Rdynload.h

typedef struct {
    const char *name;
    DL_FUNC     fun;
    int         numArgs;
} R_CallMethodDef;
typedef struct _DllInfo DllInfo;

typedef R_CallMethodDef R_ExternalMethodDef;
EXTERN DllInfo* (*R_getEmbeddingDllInfo)(void);
EXTERN int (*R_registerRoutines)(DllInfo*, void*, void*, void*, void*);


// end cdef


#ifdef _WIN32
EXTERN char *(*get_R_HOME)(void);
EXTERN char *(*getRUser)(void);
EXTERN int* UserBreak_t;
EXTERN int* CharacterMode_t;
EXTERN int* EmitEmbeddedUTF8_t;
EXTERN int (*GA_peekevent)(void);
EXTERN int (*GA_initapp)(int, char **);
#else
// eventloop.h
EXTERN void* R_InputHandlers;
EXTERN void (**R_PolledEvents_t)(void);

EXTERN void* (*R_checkActivity)(int usec, int ignore_stdin);
EXTERN void (*R_runHandlers)(void* handlers, void* mask);

EXTERN int* R_interrupts_pending_t;
#endif

#endif /* end of include guard: R_H__ */
