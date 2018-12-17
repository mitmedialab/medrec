package remoteRPC

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/mitmedialab/medrec/DatabaseManager/common"
)

type AddPermissionArgs struct {
	AgentID      string
	ViewerGroup  string
	Name         string
	StartTime    int64
	DurationDays int64
}

type AddPermissionReply struct {
	Error string
}

func (client *MedRecRemote) AddPermission(r *http.Request, args *AddPermissionArgs, reply *AddPermissionReply) error {
	log.Println("adding new permission" + args.ViewerGroup + " " + args.Name)

	tab := common.InstantiateLookupTable()
	defer tab.Close()

	rawPerms, err := tab.Get([]byte("perm-list-"+args.AgentID+"-"+args.ViewerGroup), nil)
	var perms []string
	json.Unmarshal(rawPerms, &perms)
	perms = append(perms, args.Name)
	rawPerms, err = json.Marshal(perms)
	err = tab.Put([]byte("perm-list-"+args.AgentID+"-"+args.ViewerGroup), rawPerms, nil)
	if err != nil {
		log.Println(err)
	}

	err = tab.Put([]byte("perm-startTime-"+args.AgentID+"-"+args.ViewerGroup+"-"+args.Name), []byte(strconv.FormatInt(args.StartTime, 10)), nil)
	err = tab.Put([]byte("perm-durationDays-"+args.AgentID+"-"+args.ViewerGroup+"-"+args.Name), []byte(strconv.FormatInt(args.DurationDays, 10)), nil)

	return err
}

type RemovePermissionArgs struct {
	AgentID     string
	ViewerGroup string
	Index       uint
}

type RemovePermissionReply struct {
	Error string
}

func (client *MedRecRemote) RemovePermission(r *http.Request, args *RemovePermissionArgs, reply *RemovePermissionReply) error {
	tab := common.InstantiateLookupTable()
	defer tab.Close()

	rawPerms, err := tab.Get([]byte("perm-list-"+args.AgentID+"-"+args.ViewerGroup), nil)
	var perms []string
	json.Unmarshal(rawPerms, &perms)
	name := perms[args.Index]
	perms = append(perms[:args.Index], perms[args.Index+1:]...)
	rawPerms, err = json.Marshal(perms)

	err = tab.Delete([]byte("perm-startTime-"+args.AgentID+"-"+args.ViewerGroup+"-"+name), nil)
	err = tab.Delete([]byte("perm-durationDays-"+args.AgentID+"-"+args.ViewerGroup+"-"+name), nil)

	return err
}

func (client *MedRecRemote) SetPermissionDuration(r *http.Request, args *AddPermissionArgs, reply *AddPermissionReply) error {
	tab := common.InstantiateLookupTable()
	defer tab.Close()
	err := tab.Put([]byte("perm-durationDays-"+args.AgentID+"-"+args.ViewerGroup+"-"+args.Name), []byte(strconv.FormatInt(args.DurationDays, 10)), nil)
	return err
}

func (client *MedRecRemote) SetPermissionStartTime(r *http.Request, args *AddPermissionArgs, reply *AddPermissionReply) error {
	tab := common.InstantiateLookupTable()
	defer tab.Close()
	err := tab.Put([]byte("perm-startTime-"+args.AgentID+"-"+args.ViewerGroup+"-"+args.Name), []byte(strconv.FormatInt(args.StartTime, 10)), nil)
	return err
}

type GetPermissionsArgs struct {
	AgentID     string
	ViewerGroup string
}

type Permission struct {
	Name         string
	StartTime    int64
	DurationDays int64
}

type GetPermissionsReply struct {
	Permissions []Permission
	Error       string
}

func (client *MedRecRemote) GetPermissions(r *http.Request, args *GetPermissionsArgs, reply *GetPermissionsReply) error {
	tab := common.InstantiateLookupTable()
	defer tab.Close()

	rawPerms, _ := tab.Get([]byte("perm-list-"+args.AgentID+"-"+args.ViewerGroup), nil)
	var perms []string
	json.Unmarshal(rawPerms, &perms)
	for _, name := range perms {
		_startTime, _ := tab.Get([]byte("perm-startTime-"+args.AgentID+"-"+args.ViewerGroup+"-"+name), nil)
		_durationDays, _ := tab.Get([]byte("perm-durationDays-"+args.AgentID+"-"+args.ViewerGroup+"-"+name), nil)
		startTime, _ := strconv.ParseInt(string(_startTime), 10, 64)
		durationDays, _ := strconv.ParseInt(string(_durationDays), 10, 64)

		perm := Permission{name, startTime, durationDays}
		reply.Permissions = append(reply.Permissions, perm)
	}
	return nil
}

type CheckPermissionArgs struct {
	AgentID     string
	ViewerGroup string
	Index       int
}

type CheckPermissionReply struct {
	Approved bool
	Error    string
}

//utility function for checking the access permissions of a particular viewer
func (client *MedRecRemote) CheckPermission(r *http.Request, args *CheckPermissionArgs, reply *CheckPermissionReply) error {
	tab := common.InstantiateLookupTable()
	defer tab.Close()

	rawPerms, err := tab.Get([]byte("perm-list-"+args.AgentID+"-"+args.ViewerGroup), nil)
	var perms []string
	json.Unmarshal(rawPerms, &perms)
	name := perms[args.Index]

	if len(perms) <= args.Index {
		reply.Approved = false
		return nil
	}

	_startTime, err := tab.Get([]byte("perm-startTime-"+args.AgentID+"-"+args.ViewerGroup+"-"+name), nil)
	_durationDays, err := tab.Get([]byte("perm-durationDays-"+args.AgentID+"-"+args.ViewerGroup+"-"+name), nil)
	startTime, _ := strconv.ParseInt(string(_startTime), 10, 64)
	durationDays, _ := strconv.ParseInt(string(_durationDays), 10, 32)
	if durationDays == 0 {
		reply.Approved = false
		return nil
	}
	if durationDays < 0 {
		reply.Approved = true
		return nil
	}

	//need to worry about overflow somewhat, too many durationDays and the provider will
	// won't get treated properly, but the failure mode for too many days is restricted access
	// which if you think about it is the appropriate action
	expirationTime := time.Unix(startTime, 0).AddDate(0, 0, int(durationDays))
	if expirationTime.After(time.Now()) {
		reply.Approved = true
		return nil
	}

	reply.Approved = false
	return err
}
