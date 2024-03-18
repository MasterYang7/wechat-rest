import { Component, OnInit } from '@angular/core';
import { Router, ActivatedRoute } from '@angular/router';

import { UserLevels } from '../../openapi/const';
import { RobotApi, TablesLLModel, ProfileUpdateParam } from '../../openapi/wrobot';
import { WrestApi, WcfrestContactPayload } from '../../openapi/wcfrest';


@Component({
    selector: 'page-profile-update',
    templateUrl: 'update.html'
})
export class ProfileUpdateComponent implements OnInit {

    public userLevels = UserLevels;

    public llmodels: Array<TablesLLModel> = [];

    public wcfAvatars: Record<string, string> = {};
    public wcfFriends: Array<WcfrestContactPayload> = [];
    public wcfChatrooms: Array<WcfrestContactPayload> = [];
    public wcfRoomMembers: Record<string, Array<WcfrestContactPayload>> = {};

    public conacts: Array<WcfrestContactPayload> = [];
    public conactsFilter = '';

    public formdata = {} as ProfileUpdateParam;


    constructor(
        private router: Router,
        private route: ActivatedRoute
    ) {
        this.getLLModels();
        this.getWcfFriends();
        this.getWcfChatrooms();
    }

    public ngOnInit() {
        const rd = this.route.snapshot.paramMap.get('rd');
        rd && this.getProfile(+rd);
    }

    public getProfile(rd: number) {
        RobotApi.profileDetail({ rd }).then((data) => {
            this.formdata = data;
            this.changeConacts();
        });
    }

    public updateProfile() {
        if (this.formdata.level) {
            this.formdata.level = +this.formdata.level;
        }
        RobotApi.profileUpdate(this.formdata).then(() => {
            this.router.navigate(['profile/list']);
        });
    }

    public async changeConacts() {
        const id = this.formdata.roomid || '-';
        await this.getWcfRoomMembers(this.formdata.roomid);
        this.conacts = id == '-' ? this.wcfFriends : this.wcfRoomMembers[id] || [];
    }

    public getLLModels() {
        RobotApi.llmodelList({}).then((data) => {
            this.llmodels = data || [];
        });
    }

    public getWcfFriends() {
        WrestApi.friends().then((data) => {
            this.wcfFriends = data || [];
        });
    }

    public getWcfChatrooms() {
        WrestApi.chatrooms().then((data) => {
            this.wcfChatrooms = data || [];
        });
    }

    public getWcfRoomMembers(id: string) {
        if (this.wcfRoomMembers[id]) {
            return; //已获取
        }
        return WrestApi.chatroomMembers({ roomid: id }).then((data) => {
            this.wcfRoomMembers[id] = data || [];
        });
    }

}
