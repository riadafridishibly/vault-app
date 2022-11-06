import { TablerIcon } from '@tabler/icons';
import { RouteType } from '../enums/route-type.enum';

export interface RouteData {
    path: string;
    component: () => JSX.Element;
    label?: string;
    icon: TablerIcon;
    routeType?: RouteType;
}
