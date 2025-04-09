-- M4-RO 09-OUT-23
-- FILE: T0040_084.sql
SET TIME ON;
SET DEFINE OFF;
SET PAGESIZE 20000;
SET LINESIZE 256;
SET ECHO ON;
SET SERVEROUTPUT ON;

INSERT INTO PROV_Q (SRV_TRX_S_NO , SYS_CREATION_DATE, SYS_UPDATE_DATE , OPERATOR_ID, APPLICATION_ID, DL_SERVICE_CODE, DL_UPDATE_STAMP, PRIM_RESOURCE_VAL , PRIM_RES_PARAM_CD, MARKET_CODE , PRIORITY, IGN_SUB_DEPEND_IND) VALUES  (recorder.SEQUENCE, SYSDATE, NULL, 60001, NULL,'CS030', NULL,'14994601951','MSISDN' ,'MTA', NULL,'N'); 
INSERT INTO PROV_Q_FTR (SRV_TRX_S_NO, NEW_OR_PREV, FEATURES) VALUES (recorder.SEQUENCE,'N','FTRCD=TCNAM;FTR_STATUS=A;TC_NAME=T0040_084.sql;');
INSERT INTO PROV_Q_FTR (SRV_TRX_S_NO, NEW_OR_PREV, FEATURES) VALUES (recorder.SEQUENCE,'N','FTRCD=TCDSC;FTR_STATUS=A;DESC=ESIM MNC05 MSIN36 CN14 HLRMIG;');
INSERT INTO PROV_Q_FTR (SRV_TRX_S_NO, NEW_OR_PREV, FEATURES) VALUES (recorder.SEQUENCE,'N','CUSTOMER=452669220;SUBSCRIBER_STATUS=A;PRICE_PLAN=380465367;ACCOUNT_ID=657905323;');
INSERT INTO PROV_Q_FTR (SRV_TRX_S_NO, NEW_OR_PREV, FEATURES) VALUES (recorder.SEQUENCE,'N','FTRCD=BLCYC;FTR_STATUS=A;BILLCYC=9;BLDAY=22;');
INSERT INTO PROV_Q_FTR (SRV_TRX_S_NO, NEW_OR_PREV, FEATURES) VALUES (recorder.SEQUENCE,'N','FTRCD=BOIC;FTR_STATUS=A;');
INSERT INTO PROV_Q_FTR (SRV_TRX_S_NO, NEW_OR_PREV, FEATURES) VALUES (recorder.SEQUENCE,'N','FTRCD=HLRMIG;FTR_STATUS=A;');
INSERT INTO PROV_Q_FTR (SRV_TRX_S_NO, NEW_OR_PREV, FEATURES) VALUES (recorder.SEQUENCE,'N','FTRCD=BROAI;FTR_STATUS=A;');
INSERT INTO PROV_Q_FTR (SRV_TRX_S_NO, NEW_OR_PREV, FEATURES) VALUES (recorder.SEQUENCE,'N','FTRCD=B900;FTR_STATUS=A;');
INSERT INTO PROV_Q_FTR (SRV_TRX_S_NO, NEW_OR_PREV, FEATURES) VALUES (recorder.SEQUENCE,'N','FTRCD=CAW;FTR_STATUS=A;');
INSERT INTO PROV_Q_FTR (SRV_TRX_S_NO, NEW_OR_PREV, FEATURES) VALUES (recorder.SEQUENCE,'N','FTRCD=CFU;FTR_STATUS=A;');
INSERT INTO PROV_Q_FTR (SRV_TRX_S_NO, NEW_OR_PREV, FEATURES) VALUES (recorder.SEQUENCE,'N','FTRCD=CLIID;FTR_STATUS=A;CLIENT_ID= ;');
INSERT INTO PROV_Q_FTR (SRV_TRX_S_NO, NEW_OR_PREV, FEATURES) VALUES (recorder.SEQUENCE,'N','FTRCD=CLIP;FTR_STATUS=A;');
INSERT INTO PROV_Q_FTR (SRV_TRX_S_NO, NEW_OR_PREV, FEATURES) VALUES (recorder.SEQUENCE,'N','FTRCD=CMPCD;FTR_STATUS=A;COMPANY_CODE=AC;');
INSERT INTO PROV_Q_FTR (SRV_TRX_S_NO, NEW_OR_PREV, FEATURES) VALUES (recorder.SEQUENCE,'N','FTRCD=CNPJ;FTR_STATUS=A;CNPJ_VAL=07265466000168;');
INSERT INTO PROV_Q_FTR (SRV_TRX_S_NO, NEW_OR_PREV, FEATURES) VALUES (recorder.SEQUENCE,'N','FTRCD=CPF;FTR_STATUS=A;CPF_VAL=           072;');
INSERT INTO PROV_Q_FTR (SRV_TRX_S_NO, NEW_OR_PREV, FEATURES) VALUES (recorder.SEQUENCE,'N','FTRCD=CSD;FTR_STATUS=A;');
INSERT INTO PROV_Q_FTR (SRV_TRX_S_NO, NEW_OR_PREV, FEATURES) VALUES (recorder.SEQUENCE,'N','FTRCD=CSPR;FTR_STATUS=A;');
INSERT INTO PROV_Q_FTR (SRV_TRX_S_NO, NEW_OR_PREV, FEATURES) VALUES (recorder.SEQUENCE,'N','FTRCD=CUSTN;FTR_STATUS=A;CUSTNM=CONSIMAQ COM SERV SIST DE INF EQUIP ESCRIT LTDA ME;');
INSERT INTO PROV_Q_FTR (SRV_TRX_S_NO, NEW_OR_PREV, FEATURES) VALUES (recorder.SEQUENCE,'N','FTRCD=CUSTT;FTR_STATUS=A;CUSTOMER_TYPE=B;');
INSERT INTO PROV_Q_FTR (SRV_TRX_S_NO, NEW_OR_PREV, FEATURES) VALUES (recorder.SEQUENCE,'N','FTRCD=DNLD;FTR_STATUS=A;EQOSID=308;');
INSERT INTO PROV_Q_FTR (SRV_TRX_S_NO, NEW_OR_PREV, FEATURES) VALUES (recorder.SEQUENCE,'N','FTRCD=FRTXT;FTR_STATUS=A;FREE_TEXT= ;');
INSERT INTO PROV_Q_FTR (SRV_TRX_S_NO, NEW_OR_PREV, FEATURES) VALUES (recorder.SEQUENCE,'N','FTRCD=GPRS;FTR_STATUS=A;EQOSID=314;');
INSERT INTO PROV_Q_FTR (SRV_TRX_S_NO, NEW_OR_PREV, FEATURES) VALUES (recorder.SEQUENCE,'N','FTRCD=GPRS30;FTR_STATUS=A;CODGOL=170700037;CODPCT=201608263105;FLGRC=1;');
INSERT INTO PROV_Q_FTR (SRV_TRX_S_NO, NEW_OR_PREV, FEATURES) VALUES (recorder.SEQUENCE,'N','FTRCD=HSS;FTR_STATUS=A;');
INSERT INTO PROV_Q_FTR (SRV_TRX_S_NO, NEW_OR_PREV, FEATURES) VALUES (recorder.SEQUENCE,'N','FTRCD=LMTRM;FTR_STATUS=A;');
INSERT INTO PROV_Q_FTR (SRV_TRX_S_NO, NEW_OR_PREV, FEATURES) VALUES (recorder.SEQUENCE,'N','FTRCD=MKTCD;FTR_STATUS=A;MARKET_CODE=MTA;');
INSERT INTO PROV_Q_FTR (SRV_TRX_S_NO, NEW_OR_PREV, FEATURES) VALUES (recorder.SEQUENCE,'N','FTRCD=MMS;FTR_STATUS=A;EQOSID=305;');
INSERT INTO PROV_Q_FTR (SRV_TRX_S_NO, NEW_OR_PREV, FEATURES) VALUES (recorder.SEQUENCE,'N','FTRCD=MSE152;FTR_STATUS=A;CODVAS=BDLIGHT;');
INSERT INTO PROV_Q_FTR (SRV_TRX_S_NO, NEW_OR_PREV, FEATURES) VALUES (recorder.SEQUENCE,'N','FTRCD=MSISDN;FTR_STATUS=A;MSISDN=14994601951;');
INSERT INTO PROV_Q_FTR (SRV_TRX_S_NO, NEW_OR_PREV, FEATURES) VALUES (recorder.SEQUENCE,'N','FTRCD=PA0060;FTR_STATUS=A;CODGOL=170700076;CODPCT=201609203220;FLGRC=1;');
INSERT INTO PROV_Q_FTR (SRV_TRX_S_NO, NEW_OR_PREV, FEATURES) VALUES (recorder.SEQUENCE,'N','FTRCD=PRODT;FTR_STATUS=A;PRODUCT_TYPE=Y;');
INSERT INTO PROV_Q_FTR (SRV_TRX_S_NO, NEW_OR_PREV, FEATURES) VALUES (recorder.SEQUENCE,'N','FTRCD=PRTY;FTR_STATUS=A;PRIORITY=0;');
INSERT INTO PROV_Q_FTR (SRV_TRX_S_NO, NEW_OR_PREV, FEATURES) VALUES (recorder.SEQUENCE,'N','FTRCD=PVS006;FTR_STATUS=A;CODPCT=414373146;FLGRC=1;');
INSERT INTO PROV_Q_FTR (SRV_TRX_S_NO, NEW_OR_PREV, FEATURES) VALUES (recorder.SEQUENCE,'N','FTRCD=PYMCA;FTR_STATUS=A;PAYMENT_CATEGORY=POST;');
INSERT INTO PROV_Q_FTR (SRV_TRX_S_NO, NEW_OR_PREV, FEATURES) VALUES (recorder.SEQUENCE,'N','FTRCD=RCAL;FTR_STATUS=A;');
INSERT INTO PROV_Q_FTR (SRV_TRX_S_NO, NEW_OR_PREV, FEATURES) VALUES (recorder.SEQUENCE,'N','FTRCD=RDES;FTR_STATUS=A;');
INSERT INTO PROV_Q_FTR (SRV_TRX_S_NO, NEW_OR_PREV, FEATURES) VALUES (recorder.SEQUENCE,'N','FTRCD=RSNCD;FTR_STATUS=A;REASON_CODE=OMSSP;');
INSERT INTO PROV_Q_FTR (SRV_TRX_S_NO, NEW_OR_PREV, FEATURES) VALUES (recorder.SEQUENCE,'N','FTRCD=SCHAR;FTR_STATUS=A;SCHAR=11;');
INSERT INTO PROV_Q_FTR (SRV_TRX_S_NO, NEW_OR_PREV, FEATURES) VALUES (recorder.SEQUENCE,'N','FTRCD=SIMCARD;FTR_STATUS=A;ICCID=89550531360027134321;IMSI=724053624877806;CHV5=8AD37A794E1ACFB9;KI=N44yNKy5qyHqguz3nGIJD5D5qJuJIgM4uM;PIN_1=3636;PIN_2=6005;PUK_1=84574811;PUK_2=49891761;TRANSK=24;');
INSERT INTO PROV_Q_FTR (SRV_TRX_S_NO, NEW_OR_PREV, FEATURES) VALUES (recorder.SEQUENCE,'N','FTRCD=SMO;FTR_STATUS=A;');
INSERT INTO PROV_Q_FTR (SRV_TRX_S_NO, NEW_OR_PREV, FEATURES) VALUES (recorder.SEQUENCE,'N','FTRCD=SMS;FTR_STATUS=A;');
INSERT INTO PROV_Q_FTR (SRV_TRX_S_NO, NEW_OR_PREV, FEATURES) VALUES (recorder.SEQUENCE,'N','FTRCD=SRVTR;FTR_STATUS=A;SRV_TRX_TP_CD=NEW_SUB_ACTIVATION;');
INSERT INTO PROV_Q_FTR (SRV_TRX_S_NO, NEW_OR_PREV, FEATURES) VALUES (recorder.SEQUENCE,'N','FTRCD=SUBNM;FTR_STATUS=A;SUBSCRIBER_NAME=CONSIMAQ COM SERV SIST DE INF EQUIP ESCRIT LTDA ME;');
INSERT INTO PROV_Q_FTR (SRV_TRX_S_NO, NEW_OR_PREV, FEATURES) VALUES (recorder.SEQUENCE,'N','FTRCD=SUBNO;FTR_STATUS=A;SUBSCRIBER_NO=228753096;');
INSERT INTO PROV_Q_FTR (SRV_TRX_S_NO, NEW_OR_PREV, FEATURES) VALUES (recorder.SEQUENCE,'N','FTRCD=SUBTP;FTR_STATUS=A;SUBSCRIBER_TYPE=PG3G;');
INSERT INTO PROV_Q_FTR (SRV_TRX_S_NO, NEW_OR_PREV, FEATURES) VALUES (recorder.SEQUENCE,'N','FTRCD=TWC;FTR_STATUS=A;');
INSERT INTO PROV_Q_FTR (SRV_TRX_S_NO, NEW_OR_PREV, FEATURES) VALUES (recorder.SEQUENCE,'N','FTRCD=VIDEO;FTR_STATUS=A;');
INSERT INTO PROV_Q_FTR (SRV_TRX_S_NO, NEW_OR_PREV, FEATURES) VALUES (recorder.SEQUENCE,'N','FTRCD=VM;FTR_STATUS=A;FGR=N;PROFIL=P26;');
INSERT INTO PROV_Q_FTR (SRV_TRX_S_NO, NEW_OR_PREV, FEATURES) VALUES (recorder.SEQUENCE,'N','FTRCD=VPN;FTR_STATUS=A;SUBNM=;');
INSERT INTO PROV_Q_FTR (SRV_TRX_S_NO, NEW_OR_PREV, FEATURES) VALUES (recorder.SEQUENCE,'N','FTRCD=VPNSIX;FTR_STATUS=A;');
INSERT INTO PROV_Q_FTR (SRV_TRX_S_NO, NEW_OR_PREV, FEATURES) VALUES (recorder.SEQUENCE,'N','FTRCD=VPNST;FTR_STATUS=A;CC=52422029;CCNM=CONSIMAQ COM SERV SIST DE INF EQUIP ESCR;DPTO1=52422038;DPTONM1=CONSIMAQ COM SERV SIST DE INF EQUIP ESCRCONSIMAQ COM SERV SIST DE INF EQUIP ESCR;');
INSERT INTO PROV_Q_FTR (SRV_TRX_S_NO, NEW_OR_PREV, FEATURES) VALUES (recorder.SEQUENCE,'N','FTRCD=VPNV;FTR_STATUS=A;SUBNM=;');
INSERT INTO PROV_Q_FTR (SRV_TRX_S_NO, NEW_OR_PREV, FEATURES) VALUES (recorder.SEQUENCE,'N','FTRCD=VSTRM;FTR_STATUS=A;EQOSID=308;');
INSERT INTO PROV_Q_FTR (SRV_TRX_S_NO, NEW_OR_PREV, FEATURES) VALUES (recorder.SEQUENCE,'N','FTRCD=WAP;FTR_STATUS=A;EQOSID=304;');
COMMIT;
