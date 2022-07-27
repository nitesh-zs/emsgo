/* 
 * Copyright (c) 2012-2021 TIBCO Software Inc. 
 * All Rights Reserved. Confidential & Proprietary.
 * For more information, please contact:
 * TIBCO Software Inc., Palo Alto, California, USA
 * 
 * $Id: tibufo.h 135296 2021-07-23 20:12:20Z $
 * 
 */

#ifndef _INCLUDED_tibufo_h
#define _INCLUDED_tibufo_h

#include "types.h"
#include "status.h"

#if defined(__cplusplus)
extern "C" {
#endif

extern tibemsConnectionFactory
tibemsUFOConnectionFactory_Create(void);

extern tibemsConnectionFactory
tibemsUFOConnectionFactory_CreateFromConnectionFactory(
    tibemsConnectionFactory             emsFactory);

extern tibems_status
tibemsUFOConnectionFactory_RecoverConnection(
    tibemsConnectionFactory             factory,
    tibemsConnection                    ufoConnection);

#ifdef  __cplusplus
}
#endif

#endif /* _INCLUDED_tibufo_h */
