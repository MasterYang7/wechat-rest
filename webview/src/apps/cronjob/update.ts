import { Component, OnInit } from '@angular/core';
import { Router, ActivatedRoute } from '@angular/router';

import { CronjobTypes } from '../../openapi/const';
import { SundryApi, CronjobUpdateParam } from '../../openapi/sundry';
import { WrestApi, WcfrestContactPayload } from '../../openapi/wcfrest';


@Component({
    selector: 'page-cronjob-update',
    templateUrl: 'update.html'
})
export class CronjobUpdateComponent implements OnInit {

    public cronjobTypes = CronjobTypes;

    public wcfFriends: Array<WcfrestContactPayload> = [];
    public wcfChatrooms: Array<WcfrestContactPayload> = [];
    public wcfRoomMembers: Record<string, Array<WcfrestContactPayload>> = {};

    public conacts: Array<WcfrestContactPayload> = [];
    public conactsFilter = '';

    public deliver = ['wechat', '-', '-'];

    public formdata = {} as CronjobUpdateParam;

    constructor(
        private router: Router,
        private route: ActivatedRoute
    ) {
        this.getWcfFriends();
        this.getWcfChatrooms();
    }

    public ngOnInit() {
        const rd = this.route.snapshot.paramMap.get('rd');
        rd && this.getCronjob(+rd);
    }

    public getCronjob(rd: number) {
        SundryApi.cronjobDetail({ rd }).then((data) => {
            this.formdata = data;
            const dataDeliver = data.deliver.split(',');
            for (const [k, v] of dataDeliver.entries()) {
                this.deliver[k] = v;
            }
            this.changeConacts();
        });
    }

    public updateCronjob() {
        const data = this.formdata;
        const time = data.second + data.minute + data.hour + data.day_of_month + data.month + data.day_of_week;
        if (time === '******') {
            window.postMessage({ message: '排程不可全为 *', type: 'danger' });
            return;
        }
        this.formdata.deliver = Object.values(this.deliver).join(',');
        SundryApi.cronjobUpdate(this.formdata).then(() => {
            this.router.navigate(['cronjob/list']);
        });
    }

    public async changeConacts() {
        const id = this.deliver[1] || '-';
        await this.getWcfRoomMembers(this.deliver[1]);
        this.conacts = id == '-' ? this.wcfFriends : this.wcfRoomMembers[id] || [];
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
