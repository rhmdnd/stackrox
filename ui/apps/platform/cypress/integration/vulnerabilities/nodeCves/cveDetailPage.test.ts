import withAuth from '../../../helpers/basicAuth';
import { hasFeatureFlag } from '../../../helpers/features';

// TODO - classic tests conditionally skip many of these checks when running on OpenShift, we
//        need to verify if this is still necessary and if so, how to handle it in the new tests
describe('Node CVEs - CVE Detail Page', () => {
    withAuth();

    before(function () {
        if (!hasFeatureFlag('ROX_VULN_MGMT_NODE_PLATFORM_CVES')) {
            this.skip();
        }
    });

    it('should restrict access to users with insufficient permissions', () => {
        // check that users without Node access cannot access the Node CVE Detail page directly
    });

    it('should only show relevant filters for the page', () => {
        // check the advanced filters and ensure only the relevant filters are displayed
    });

    it('should link to the overview page from the breadcrumbs', () => {
        // clicking the Node CVEs breadcrumb should navigate to the overview page with the CVE tab selected
    });

    it('should link to the Node page from the name links in the table', () => {
        // clicking a Node name in the list should navigate to the correct Node details page
    });

    it('should display the expected Node table columns', () => {
        // check presence of Node column
        // check presence of CVE Severity column
        // check presence of CVE status column
        // check presence of CVSS column
        // check presence of cluster column
        // check presence of Operating System column
        // check presence of Affected components column
    });

    it('should sort Node table columns', () => {
        // check sorting of Node column
        // check sorting of CVE Severity column
        // check sorting of CVE status column
        // check sorting of CVSS column
        // check sorting of cluster column
        // check sorting of Operating System column
        // check sorting of Affected components column
    });

    it('should filter the Node table', () => {
        // filtering by Node name should only display rows with a matching name
        // filtering by CVE Severity should only display rows with a matching severity
        // filtering by CVE Status should only display rows with a matching status
        // filtering by CVSS should only display rows with a CVSS in range
        // filtering by Cluster should only display rows with a matching cluster
        // filtering by Operating System should only display rows with a matching OS
        // filtering by Component should only display rows with a nested table containing a matching component
        //   - expand each row
        //   - check that the component name exists in the table
        // clearing filters should remove all filter chips and filter from the URL
    });

    // Note: This might not be reliable in CI due to the low number of Nodes. We may need to mock and/or
    // just test the logic of the pagination in the URL.
    it('should paginate the Node table', () => {
        // visit the location and save the list of Node names with a default perPage
        //   visit the location with perPage=2 in the URL
        //     should only display the first 2 rows of the previous list
        //     paginating to the next page should display the following two rows
        // go to page 1, then page 2
        //   applying a filter should reset the page to 1
        // go to page 1, then page 2
        //   go to next page, applying a sort should reset the page to 1
    });

    it('should update summary cards when a filter is applied', () => {
        // apply a Critical severity filter and ensure that Important/Moderate/Low severities read "Results hidden" in the card
        // clear filters
        // apply a CVE status filter and ensure that the opposite status reads "Results hidden" in the card
    });

    it('should allow viewing the Node details', () => {
        // click the Details tab and ensure that the vulnerabilities table no longer displays
        // verify that a Cluster field exists
        // verify that label and annotation sections exist
    });
});
