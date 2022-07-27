
/* 
 * Copyright (c) 2001-2021 TIBCO Software Inc. 
 * All Rights Reserved. Confidential & Proprietary.
 * For more information, please contact:
 * TIBCO Software Inc., Palo Alto, California, USA
 * 
 * $Id: clookup.h 136308 2021-09-08 21:38:22Z $
 * 
 */

#ifndef _INCLUDED_tibems_clookup_h
#define _INCLUDED_tibems_clookup_h

#include "types.h"
#include "status.h"
#include "dest.h"
#include "confact.h"

#if defined(__cplusplus)
extern "C" {
#endif



/*************************************************************************/
/* public types used by the API                                          */
/*************************************************************************/

typedef void*              tibemsLookupContext;

/*************************************************************************/
/* public API                                                            */
/*************************************************************************/
extern tibems_status
tibemsLookupContext_Create(
    tibemsLookupContext*    context,
    const char*             brokerURL,
    const char*             username,
    const char*             password);

/* obsolete as of release 10.1 */
extern tibems_status
tibemsLookupContext_CreateExternal(
    tibemsLookupContext*    context,
    tibemsLookupParams      lookupParams);

extern tibems_status
tibemsLookupContext_CreateSSL(
    tibemsLookupContext*    context,
    const char*             brokerURL,
    const char*             username,
    const char*             password,
    tibemsSSLParams         SSLparams,
    const char*             pk_password);

extern tibems_status
tibemsLookupContext_Destroy(
    tibemsLookupContext     context);
    

/* generic lookup any object type */
extern tibems_status
tibemsLookupContext_Lookup(
    tibemsLookupContext         context,
    const char*                 name,
    void**                      object);

/* lookup destination */
extern tibems_status
tibemsLookupContext_LookupDestination(
    tibemsLookupContext         context,
    const char*                 name,
    tibemsDestination*          destination);

/* lookup connection factory */
extern tibems_status
tibemsLookupContext_LookupConnectionFactory(
    tibemsLookupContext         context,
    const char*                 name,
    tibemsConnectionFactory*    factory);


/* obsolete as of release 10.1 */
extern tibemsLookupParams
tibemsLookupParams_Create(void);

/* obsolete as of release 10.1 */
extern void
tibemsLookupParams_Destroy(
    tibemsLookupParams      lparams);

/* obsolete as of release 10.1 */
extern tibems_status
tibemsLookupParams_SetLdapServerUrl(
    tibemsLookupParams      lparams,
    const char*             url);

/* obsolete as of release 10.1 */
extern char*
tibemsLookupParams_GetLdapServerUrl(
    tibemsLookupParams      lparams);

/* obsolete as of release 10.1 */
extern tibems_status
tibemsLookupParams_SetLdapBaseDN(
    tibemsLookupParams      lparams,
    const char*             basedn);

/* obsolete as of release 10.1 */
extern tibems_status
tibemsLookupParams_SetLdapPrincipal(
    tibemsLookupParams      lparams,
    const char*             principal);

/* obsolete as of release 10.1 */
extern tibems_status
tibemsLookupParams_SetLdapCredential(
    tibemsLookupParams      lparams,
    const char*             credential);

/* obsolete as of release 10.1 */
extern tibems_status
tibemsLookupParams_SetLdapSearchScope(
    tibemsLookupParams      lparams,
    const char*             scope);

/* obsolete as of release 10.1 */
extern tibems_status
tibemsLookupParams_SetLdapConnType(
    tibemsLookupParams      lparams,
    const char*             type);

/* obsolete as of release 10.1 */
extern tibems_status
tibemsLookupParams_SetLdapCAFile(
    tibemsLookupParams      lparams,
    const char*             file);

/* obsolete as of release 10.1 */
extern tibems_status
tibemsLookupParams_SetLdapCAPath(
    tibemsLookupParams      lparams,
    const char*             path);

/* obsolete as of release 10.1 */
extern tibems_status
tibemsLookupParams_SetLdapCertFile(
    tibemsLookupParams      lparams,
    const char*             file);

/* obsolete as of release 10.1 */
extern tibems_status
tibemsLookupParams_SetLdapKeyFile(
    tibemsLookupParams      lparams,
    const char*             file);

/* obsolete as of release 10.1 */
extern tibems_status
tibemsLookupParams_SetLdapRandFile(
    tibemsLookupParams      lparams,
    const char*             file);

/* obsolete as of release 10.1 */
extern tibems_status
tibemsLookupParams_SetLdapCiphers(
    tibemsLookupParams      lparams,
    const char*             ciphers);

#ifdef  __cplusplus
}
#endif

#endif /* _INCLUDED_tibems_clookup_h */

