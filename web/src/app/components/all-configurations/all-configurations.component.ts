import {Component, ElementRef, OnInit, ViewChild} from '@angular/core';
import {DataSource} from '@angular/cdk/collections';
import {BehaviorSubject, Observable} from 'rxjs/Rx';
import {MatPaginator, MatSort} from '@angular/material';
import {TransactionDataService} from '../../services/transaction.data.service';
import {HttpClient} from '@angular/common/http';
import {HttpService} from '../../services/http.service';
import {Transaction} from '../../models/transaction';
import {animate, state, style, transition, trigger} from '@angular/animations';

@Component({
  selector: 'app-all-configurations',
  templateUrl: './all-configurations.component.html',
  styleUrls: ['./all-configurations.component.scss'],
  animations: [
    trigger('detailExpand', [
      state('collapsed', style({height: '0px', minHeight: '0', display: 'none'})),
      state('expanded', style({height: '*'})),
      transition('expanded <=> collapsed', animate('225ms cubic-bezier(0.4, 0.0, 0.2, 1)')),
    ]),
  ],
})
export class AllConfigurationsComponent implements OnInit {
  dataSource: ExampleDataSource | null;
  columnsToDisplay = ['hash', 'createdAt', 'updatedAt', 'chainType'];
  expandedElement: Transaction;

  constructor(public httpClient: HttpClient, public httpService: HttpService, public transactionDataService: TransactionDataService) { }

  @ViewChild(MatPaginator) paginator: MatPaginator;
  @ViewChild(MatSort) sort: MatSort;
  @ViewChild('filter') filter: ElementRef;

  ngOnInit() {
    this.loadData(this);
  }

  public loadData(_this) {
    _this.dataService = new TransactionDataService(_this.httpClient);
    _this.dataSource = new ExampleDataSource(_this.dataService, _this.paginator, _this.sort);
    Observable.fromEvent(_this.filter.nativeElement, 'keyup')
      .debounceTime(150)
      .distinctUntilChanged()
      .subscribe(() => {
        if (!_this.dataSource) {
          return;
        }
        let stringValue = _this.filter.nativeElement.value.toLowerCase();
        _this.dataSource.filter = stringValue;
      });
  }
}


export class ExampleDataSource extends DataSource<any> {
  _filterChange = new BehaviorSubject('');
  _filterStatus = new BehaviorSubject('');

  get filter(): string {
    return this._filterChange.value;
  }

  set filter(filter: string) {
    this._filterChange.next(filter);
  }

  get activeState(): string {
    return this._filterStatus.value;
  }

  set activeState(activeState: string) {
    this._filterStatus.next(activeState);
  }

  filteredData: any[] = [];
  renderedData: any[] = [];

  constructor(public _exampleDatabase: TransactionDataService,
              public _paginator: MatPaginator,
              public _sort: MatSort) {
    super();
    // Reset to the first page when the user changes the filter.
    this._filterChange.subscribe(() => this._paginator.pageIndex = 0);
  }

  /** Connect function called by the table to retrieve one stream containing the data to render. */
  connect(): Observable<Transaction[]> {
    // Listen for any changes in the base data, sorting, filtering, or pagination
    const displayDataChanges = [
      this._exampleDatabase.dataChange,
      this._sort.sortChange,
      this._filterChange,
      this._filterStatus,
      this._paginator.page
    ];

    this._exampleDatabase.getAllTransactions();

    return Observable.merge(...displayDataChanges).map(() => {
      // Filter data
      this.filteredData = this._exampleDatabase.data.slice().filter((transaction: Transaction) => {
        const searchStr = (transaction.id + transaction.hash).toLowerCase();
        return searchStr.indexOf(this.filter.toLowerCase()) !== -1;
      });

      // Sort filtered data
      const sortedData = this.sortData(this.filteredData.slice());

      // Grab the page's slice of the filtered sorted data.
      const startIndex = this._paginator.pageIndex * this._paginator.pageSize;
      this.renderedData = sortedData.splice(startIndex, this._paginator.pageSize);
      return this.renderedData;
    });
  }
  disconnect() {
  }



  /** Returns a sorted copy of the database data. */
  sortData(data: any[]): any[] {
    if (!this._sort.active || this._sort.direction === '') {
      return data;
    }

    return data.sort((a, b) => {
      let propertyA: number | string = '';
      let propertyB: number | string = '';

      switch (this._sort.active) {
        case 'id': [propertyA, propertyB] = [a.id, b.id]; break;
        case 'username': [propertyA, propertyB] = [a.username, b.username]; break;
        case 'createdAt': [propertyA, propertyB] = [a.createdAt, b.createdAt]; break;
        case 'status': [propertyA, propertyB] = [a.status, b.status]; break;
        case 'skycoinAddress': [propertyA, propertyB] = [a.skycoinAddress, b.skycoinAddress]; break;
      }

      const valueA = isNaN(+propertyA) ? propertyA : +propertyA;
      const valueB = isNaN(+propertyB) ? propertyB : +propertyB;

      return (valueA < valueB ? -1 : 1) * (this._sort.direction === 'asc' ? 1 : -1);
    });
  }
}

