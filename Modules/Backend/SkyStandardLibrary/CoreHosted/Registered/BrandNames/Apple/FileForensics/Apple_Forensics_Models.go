//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//	This file contains all information for the SkyLine standard library //   _____ _       __    _              _____         _
//	such as the math, io, IoT, ECP, Socket, Web, HTTP and other various //  |   __| |_ _ _|  |  |_|___ ___ ___ |  _  |___ ___| |___
//	library content types. This code section defines a sub unit under a //  |__   | '_| | |  |__| |   | -_|___ |     | . | . | | -_|
//	primary unit. This means that this section is under neath all of the//  |_____|_,_|_  |_____|_|_|_|___|    |__|__|  _|  _|_|___|
//	primary units that define the functions to register the sub func's  //            |___|                          |_| |_|
//////////////////////////////////////////////////////////////////////////
//||||||||||||||||||||||||||
// Apple Library Unoffical | This library is defined specifically for Apple related functions, creations, files, protocols, programs, exploits and much more. This means that
//||||||||||||||||||||||||||
//
// anything within this library will be used for Apple related tasks only. The reason this is under "Brands" is because Apple has their own unique frameworks, protocols, systems
//
// code, tools, backends, language's and so much more along that list. The tools and utilities within this program have ZERO INTENTION OF CAUSING HARM AS THE USE OF THE PROGRAM IS NOT
//
// DEVELOPED TO REVERSE ENGINEER, AUTOMATE EXPLOITS, AUTOMATE TERMS AND ISSUES FOR PROTOCOLS AND PROGRAMS USED WITHIN THESE DEVICES AND THE CREATORS OF THE SKYLINE PROGRAMMING LANGUAGE
//
// ARE NOT HELD LIABLE FOR ANY MISUSE OF THESE TOOLS. THESE TOOLS ARE DESIGNED FOR EDUCATIONAL PURPOSED AND SHOULD BE USED ACCORDINGLY, THIS MEANS THAT THE USER OF THE SOFTWARE IS USING IT
//
// FOR THE USE CASE OF EDUCATING THE GENERAL PUBLIC OR OTHERS ABOUT THE INTERNALS AND STRUCTURES OF PROTOCOLS, FILES, APPLICATIONS, DEVICES, HARDWARE, NETWORKS, BINARIES, CODES, SYSTEMS, API'S
//
// WEB TOOLS, UTILITIES, DEV TOOLS, SDK'S AND OTHER VARIOUS SOFTWARE. THIS SOFTWARE WAS DEVELOPED FOR THE INTENTION TO INFORM AND EDUCATE UPCOMING APPLE DEVELOPERS, PROGRAMMERS AND SECURITY
//
// RESEARCHERS TO BETTER UNDERSTAND HOW WELL BUILT, STRUCTURED, DEEP AND STATIC APPLE AS A PRODUCT GOES. THE CODES USED HERE IS ALL PUBLIC INFORMATION GIVEN BY APPLE AND GIVEN BY COMMUNITY MEMBERS
//
// AND DOES NOT BREAK ANY LAW TALKING ABOUT COPYWRITE. APPLE'S SOFTWARE EVEN UNDER UNOFICAL DOCUMENTATION AS WELL AS THE INFORMATION USED TO CREATE TOOLS FOR THIS FRAMEWORK WERE ASSISTED BY GIVEN
//
// EMPLOYEES AT APPLE.INC AND THE INFORMATION GATHERED WAS RELEASED ON PUBLIC DOCUMENTS SIGNED OR WRITTEN BY APPLE AS A COPANY, THE DATA HERE WAS ALSO RELEASED PUBLICLY OR IS ACCESSIBLE BY ANY OTHER
//
// DEVICE AS APART OF SERVICE DISCOVERY, INFORMATION DISCOVERY, SERVICE CONNECTION, SERVICE DEVELOPMENT, SERVICE WORKS, NETWORKS OR INTERNAL COMMUNICATIONS. INFORMATION USED WITHIN THIS LIBRARY
//
// ABOUT PROTOCOLS, API'S, SYSTEMS AND OTHER SOFTWARE DEVELOPED BY APPLE WAS OBTAINED AGAIN THROUGH THE USE OF PUBLIC INFORMATION AND UNDERSTANDING USING APPLE'S OWN SDK'S SUCH AS BONJOUR.
//
//::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
//
// File defines: All models, structures, constants, paths, files, maps etc used within the library
//
//
package Apple_ForensicsLibrary

const (
	//:::::::::::::::::::::::::::::::::::::::::::
	//:: DAAP server response codes by abrev
	//:::::::::::::::::::::::::::::::::::::::::::
	Apple_DAAP_Container_MSRV             = "msrv"
	Apple_DAAP_Status_MSTT                = "mstt"
	Apple_DAAP_ProtocolVersion_APRO       = "apro"
	Apple_DAAP_ItemName_MINM              = "minm"
	Apple_DAAP_MusicSharingVersion_AESV   = "aeSV"
	Apple_DAAP_TimeoutInterval_MSTM       = "mstm"
	Apple_DAAP_DatabasesCount_MSFC        = "msdc"
	Apple_DAAP_ReqFplay_AEFP              = "aeFP"
	Apple_DAAP_UnknownTag_AEFR            = "aeFR"
	Apple_DAAP_LoginRequired_MSLR         = "mslr"
	Apple_DAAP_SupportsAutoLogout_MSAL    = "msal"
	Apple_DAAP_UtcTime_MSTC               = "mstc"
	Apple_DAAP_UtcOffset_MSTO             = "msto"
	Apple_DAAP_UnknownTag2_ATSV           = "atSV"
	Apple_DAAP_SupportsExtraData_ATED     = "ated"
	Apple_DAAP_GaplessResy_ASGR           = "asgr"
	Apple_DAAP_UnknownTag3_ASSE           = "asse"
	Apple_DAAP_UnknownTag4_AESX           = "aeSX"
	Apple_DAAP_SupportsEdit_MSED          = "msed"
	Apple_DAAP_SupportsUpdate_MSUP        = "msup"
	Apple_DAAP_SupportsPersistentIDs_MSPI = "mspi"
	Apple_DAAP_SupportsExtensions_MSEX    = "msex"
	Apple_DAAP_SupportsBrowse_MSBR        = "msbr"
	Apple_DAAP_SupportsQuery_MSQY         = "msqy"
	Apple_DAAP_SupportsIndex_MSIX         = "msix"
	Apple_DAAP_UnknownTag5_MSCU           = "mscu"
	//::::::::::::::::::::::::::::::::::::::::::::::::
	//:: File codes, signatures, suffixes and others
	//::::::::::::::::::::::::::::::::::::::::::::::::
	Apple_BinaryPlist_Sign_V0          = "bplist00"
	Apple_BinaryPlist_Sign_V1          = "bplist01"
	Apple_BinaryPlist_Sign_Suffix      = "*.bplist"
	Apple_BinaryPlist_Sign_Description = "Apple Binary Property List Format (BPLIST)"
	Apple_BinaryPlist_Magic_Byte       = "62 70 6c 69 73 74 30 30"
	Apple_Plist_HeaderVerification     = `<?xml version=\"1.0\" encoding=\"UTF-8\"?>\n<!DOCTYPE plist PUBLIC \"-//Apple//DTD PLIST 1.0//EN\" \"http://www.apple.com/DTDs/PropertyList-1.0.dtd\">\n<plist version=\"1.0\">\n`
	//::::::::::::::::::::::::::::::::::::::::::::::::
	//:: Bplist data type tags
	//::::::::::::::::::::::::::::::::::::::::::::::::
	Apple_BinaryPlist_DataType_String_ASCII = 0x50
	Apple_BinaryPlist_DataType_String_UFT16 = 0x60
	Apple_BinaryPlist_DataType_Array        = 0xA0
	Apple_BinaryPlist_DataType_Hash         = 0xD0
	Apple_BinaryPlist_DataType_Null         = 0x00
	Apple_BinaryPlist_DataType_BooleanFalse = 0x08
	Apple_BinaryPlist_DataType_BooleanTrue  = 0x09
	Apple_BinaryPlist_DataType_Integer      = 0x10
	Apple_BinaryPlist_DataType_Real         = 0x20
	//::::::::::::::::::::::::::::::::::::::::::::::::
	//:: Bplist tags and other various information
	//::::::::::::::::::::::::::::::::::::::::::::::::
	Apple_BinaryPlist_Tag_DataTag = 0x40
	Apple_BinaryPlist_Tag_Date    = 0x30
	Apple_BinaryPlist_Tag_UID     = 0x80
)
